package producers

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nikitaNotFound/evm-indexer-go/internal/config"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/nikitaNotFound/evm-indexer-go/internal/networks"
	"github.com/nikitaNotFound/evm-indexer-go/internal/smartcontracts/abigen"
	"github.com/rs/zerolog/log"
)

const (
	UniswapV3PoolsTopicName = "uniswap_v3_pools"
)

type UniswapV3PoolsProducer struct {
	ethClient      *ethclient.Client
	factoryAddress string
	factoryAbi     *abigen.UniswapV3Factory
	cfg            *config.Config
}

func NewUniswapV3PoolsProducer(
	ethClient *ethclient.Client,
	cfg *config.Config,
) *UniswapV3PoolsProducer {
	factoryAddress := networks.GetUniswapV3FactoryAddress(cfg.NetworkConfig.Network)

	factoryAbi, err := abigen.NewUniswapV3Factory(common.HexToAddress(factoryAddress), ethClient)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create uniswap v2 factory")
	}

	return &UniswapV3PoolsProducer{
		ethClient:      ethClient,
		factoryAddress: factoryAddress,
		factoryAbi:     factoryAbi,
		cfg:            cfg,
	}
}

type UniswapV3Pool struct {
	Address     string
	Token0      string
	Token1      string
	Fee         *big.Int
	TickSpacing *big.Int
}

// TODO: refactor to use shared workpool for blocks producer
func (p *UniswapV3PoolsProducer) OnProduce(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().Str("component", "UniswapV3PoolsProducer").Str("method", "OnProduce").Logger()

	l.Info().Interface("trigger", trigger).Msg("producing blocks")

	if trigger.EndBlock-trigger.StartBlock >= 1 {
		return p.scanPools(e, trigger)
	}

	return ErrInvalidBlockRange
}

func (p *UniswapV3PoolsProducer) scanPools(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().
		Str("component", "UniswapV3PoolsProducer").
		Str("method", "handleBatchLoad").
		Logger()

	endBlock := uint64(trigger.EndBlock)

	logs, err := p.factoryAbi.FilterPoolCreated(
		&bind.FilterOpts{
			Start:   uint64(trigger.StartBlock),
			End:     &endBlock,
			Context: e.Ctx,
		},
		[]common.Address{},
		[]common.Address{},
		[]*big.Int{},
	)
	if err != nil {
		l.Error().Err(err).Msg("failed to filter pair created")
		return err
	}

	v3PoolEvents := make([]models.ProducedDataEvent, 0)
	for logs.Next() {
		log := logs.Event

		v3PoolEvents = append(v3PoolEvents, models.NewProducedDataEvent(&UniswapV3Pool{
			Address:     log.Pool.Hex(),
			Token0:      log.Token0.Hex(),
			Token1:      log.Token1.Hex(),
			Fee:         log.Fee,
			TickSpacing: log.TickSpacing,
		}))
	}

	if err := e.BroadcastData(
		e.Ctx,
		UniswapV3PoolsTopicName,
		v3PoolEvents,
	); err != nil {
		l.Error().Err(err).Msg("failed to broadcast data")
		return err
	}

	return nil
}

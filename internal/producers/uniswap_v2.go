package producers

import (
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
	UniswapV2PoolsTopicName = "uniswap_v2_pools"
)

type UniswapV2PoolsProducer struct {
	ethClient      *ethclient.Client
	factoryAddress string
	factoryAbi     *abigen.UniswapV2Factory
	cfg            *config.Config
}

func NewUniswapV2PoolsProducer(
	ethClient *ethclient.Client,
	cfg *config.Config,
) *UniswapV2PoolsProducer {
	factoryAddress := networks.GetUniswapV2FactoryAddress(cfg.NetworkConfig.Network)

	factoryAbi, err := abigen.NewUniswapV2Factory(common.HexToAddress(factoryAddress), ethClient)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create uniswap v2 factory")
	}

	return &UniswapV2PoolsProducer{
		ethClient:      ethClient,
		factoryAddress: factoryAddress,
		factoryAbi:     factoryAbi,
		cfg:            cfg,
	}
}

type UniswapV2Pool struct {
	Address string
	Token0  string
	Token1  string
}

// TODO: refactor to use shared workpool for blocks producer
func (p *UniswapV2PoolsProducer) OnProduce(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().Str("component", "UniswapV2PoolsProducer").Str("method", "OnProduce").Logger()

	l.Info().Interface("trigger", trigger).Msg("producing blocks")

	if trigger.EndBlock-trigger.StartBlock >= 1 {
		return p.scanPools(e, trigger)
	}

	return ErrInvalidBlockRange
}

func (p *UniswapV2PoolsProducer) scanPools(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().
		Str("component", "UniswapV2PoolsProducer").
		Str("method", "handleBatchLoad").
		Logger()

	endBlock := uint64(trigger.EndBlock)

	logs, err := p.factoryAbi.FilterPairCreated(
		&bind.FilterOpts{
			Start:   uint64(trigger.StartBlock),
			End:     &endBlock,
			Context: e.Ctx,
		},
		[]common.Address{},
		[]common.Address{},
	)
	if err != nil {
		l.Error().Err(err).Msg("failed to filter pair created")
		return err
	}

	v2PoolEvents := make([]models.ProducedDataEvent, 0)
	for logs.Next() {
		log := logs.Event

		v2PoolEvents = append(v2PoolEvents, models.NewProducedDataEvent(&UniswapV2Pool{
			Address: log.Pair.Hex(),
			Token0:  log.Token0.Hex(),
			Token1:  log.Token1.Hex(),
		}))
	}

	if err := e.BroadcastData(
		e.Ctx,
		UniswapV2PoolsTopicName,
		v2PoolEvents,
	); err != nil {
		l.Error().Err(err).Msg("failed to broadcast data")
		return err
	}

	return nil
}

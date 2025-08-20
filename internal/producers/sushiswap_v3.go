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
	SushiSwapV3PoolsTopicName = "sushiswap_v3_pools"
)

type SushiSwapV3PoolsProducer struct {
	ethClient      *ethclient.Client
	factoryAddress string
	factoryAbi     *abigen.UniswapV3Factory
	cfg            *config.Config
}

func NewSushiSwapV3PoolsProducer(
	ethClient *ethclient.Client,
	cfg *config.Config,
) *SushiSwapV3PoolsProducer {
	factoryAddress := networks.GetSushiSwapV3FactoryAddress(cfg.NetworkConfig.Network)

	log.Info().
		Str("factory_address", factoryAddress).
		Str("network", string(cfg.NetworkConfig.Network)).
		Msg("creating SushiSwapV3 producer")

	factoryAbi, err := abigen.NewUniswapV3Factory(common.HexToAddress(factoryAddress), ethClient)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create sushiswap v3 factory")
	}

	return &SushiSwapV3PoolsProducer{
		ethClient:      ethClient,
		factoryAddress: factoryAddress,
		factoryAbi:     factoryAbi,
		cfg:            cfg,
	}
}

type SushiSwapV3Pool struct {
	Address     string
	Token0      string
	Token1      string
	Fee         *big.Int
	TickSpacing *big.Int
}

func (p *SushiSwapV3PoolsProducer) OnProduce(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().Str("component", "SushiSwapV3PoolsProducer").Str("method", "OnProduce").Logger()

	l.Info().Interface("trigger", trigger).Msg("producing blocks")

	if trigger.EndBlock-trigger.StartBlock >= 1 {
		return p.scanPools(e, trigger)
	}

	return ErrInvalidBlockRange
}

func (p *SushiSwapV3PoolsProducer) scanPools(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().
		Str("component", "SushiSwapV3PoolsProducer").
		Str("method", "scanPools").
		Logger()

	endBlock := uint64(trigger.EndBlock)

	l.Info().
		Str("factory_address", p.factoryAddress).
		Uint64("start_block", uint64(trigger.StartBlock)).
		Uint64("end_block", endBlock).
		Msg("scanning for PoolCreated events")

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
		l.Error().Err(err).Msg("failed to filter pool created")
		return err
	}

	v3PoolEvents := make([]models.ProducedDataEvent, 0)
	for logs.Next() {
		log := logs.Event

		v3PoolEvents = append(v3PoolEvents, models.NewProducedDataEvent(&SushiSwapV3Pool{
			Address:     log.Pool.Hex(),
			Token0:      log.Token0.Hex(),
			Token1:      log.Token1.Hex(),
			Fee:         log.Fee,
			TickSpacing: log.TickSpacing,
		}))
	}

	if logs.Error() != nil {
		l.Error().Err(logs.Error()).Msg("error iterating over PoolCreated events")
		return logs.Error()
	}

	l.Info().Int("count", len(v3PoolEvents)).
		Int64("start_block", trigger.StartBlock).
		Int64("end_block", trigger.EndBlock).
		Msg("produced data events")

	if err := e.BroadcastData(
		e.Ctx,
		SushiSwapV3PoolsTopicName,
		v3PoolEvents,
	); err != nil {
		l.Error().Err(err).Msg("failed to broadcast data")
		return err
	}

	return nil
}
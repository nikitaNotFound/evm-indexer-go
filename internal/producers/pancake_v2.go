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
	PancakeV2PoolsTopicName = "pancake_v2_pools"
)

type PancakeV2PoolsProducer struct {
	ethClient      *ethclient.Client
	factoryAddress string
	factoryAbi     *abigen.UniswapV2Factory
	cfg            *config.Config
}

func NewPancakeV2PoolsProducer(
	ethClient *ethclient.Client,
	cfg *config.Config,
) *PancakeV2PoolsProducer {
	factoryAddress := networks.GetPancakeV2FactoryAddress(cfg.NetworkConfig.Network)

	log.Info().
		Str("factory_address", factoryAddress).
		Str("network", string(cfg.NetworkConfig.Network)).
		Msg("creating PancakeV2 producer")

	factoryAbi, err := abigen.NewUniswapV2Factory(common.HexToAddress(factoryAddress), ethClient)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create pancake v2 factory")
	}

	return &PancakeV2PoolsProducer{
		ethClient:      ethClient,
		factoryAddress: factoryAddress,
		factoryAbi:     factoryAbi,
		cfg:            cfg,
	}
}

type PancakeV2Pool struct {
	Address string
	Token0  string
	Token1  string
}

// TODO: refactor to use shared workpool for blocks producer
func (p *PancakeV2PoolsProducer) OnProduce(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().Str("component", "PancakeV2PoolsProducer").Str("method", "OnProduce").Logger()

	l.Info().Interface("trigger", trigger).Msg("producing blocks")

	if trigger.EndBlock-trigger.StartBlock >= 1 {
		return p.scanPools(e, trigger)
	}

	return ErrInvalidBlockRange
}

func (p *PancakeV2PoolsProducer) scanPools(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().
		Str("component", "PancakeV2PoolsProducer").
		Str("method", "handleBatchLoad").
		Logger()

	endBlock := uint64(trigger.EndBlock)

	l.Info().
		Str("factory_address", p.factoryAddress).
		Uint64("start_block", uint64(trigger.StartBlock)).
		Uint64("end_block", endBlock).
		Msg("scanning for PairCreated events")

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

		v2PoolEvents = append(v2PoolEvents, models.NewProducedDataEvent(&PancakeV2Pool{
			Address: log.Pair.Hex(),
			Token0:  log.Token0.Hex(),
			Token1:  log.Token1.Hex(),
		}))
	}
	
	// Check for iterator error
	if logs.Error() != nil {
		l.Error().Err(logs.Error()).Msg("error iterating over PairCreated events")
		return logs.Error()
	}

	l.Info().Int("count", len(v2PoolEvents)).
		Int64("start_block", trigger.StartBlock).
		Int64("end_block", trigger.EndBlock).
		Msg("produced data events")

	if err := e.BroadcastData(
		e.Ctx,
		PancakeV2PoolsTopicName,
		v2PoolEvents,
	); err != nil {
		l.Error().Err(err).Msg("failed to broadcast data")
		return err
	}

	return nil
}

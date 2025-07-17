package producers

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/nikitaNotFound/evm-indexer-go/pkg/workpool"
	"github.com/rs/zerolog/log"
)

const TopicName = "blocks"

var (
	ErrInvalidBlockRange = errors.New("invalid block range")
)

type BlocksProducer struct {
	ethClient *ethclient.Client
}

func NewBlocksProducer(ethClient *ethclient.Client) *BlocksProducer {
	return &BlocksProducer{ethClient: ethClient}
}

type Block struct {
	Number int64
	Hash   string
}

func (p *BlocksProducer) OnProduce(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().Str("component", "BlocksProducer").Str("method", "OnProduce").Logger()

	l.Info().Interface("trigger", trigger).Msg("producing blocks")

	if trigger.EndBlock-trigger.StartBlock > 1 {
		return p.handleBatchLoad(e, trigger)
	}

	if trigger.EndBlock-trigger.StartBlock == 1 {
		return p.handleSingleLoad(e, trigger)
	}

	return ErrInvalidBlockRange
}

func (p *BlocksProducer) handleBatchLoad(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().Str("component", "BlocksProducer").Str("method", "handleBatchLoad").Logger()

	// TODO: later load this amount from config
	workPool := workpool.NewWorkPool[*Block](10)
	defer workPool.Stop()

	for i := trigger.StartBlock; i <= trigger.EndBlock; i++ {
		workPool.Do(func() (*Block, error) {
			return p.loadBlockInfo(i, e.Ctx)
		})
	}

	go func() {
		for block := range workPool.Results() {
			if err := e.BroadcastData(TopicName, block); err != nil {
				l.Error().Err(err).Msg("failed to broadcast data")
			}
		}
	}()

	workPool.Wait()

	return nil
}

func (p *BlocksProducer) handleSingleLoad(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().Str("component", "BlocksProducer").Str("method", "handleSingleLoad").Logger()

	block, err := p.loadBlockInfo(trigger.StartBlock, e.Ctx)
	if err != nil {
		return err
	}

	if err := e.BroadcastData(TopicName, block); err != nil {
		l.Error().Err(err).Msg("failed to broadcast data")
		return err
	}

	return nil
}

func (p *BlocksProducer) loadBlockInfo(
	blockNumber int64,
	ctx context.Context,
) (*Block, error) {
	block, err := p.ethClient.BlockByNumber(ctx, big.NewInt(blockNumber))
	if err != nil {
		return nil, err
	}

	return &Block{Number: int64(blockNumber), Hash: block.Hash().Hex()}, nil
}

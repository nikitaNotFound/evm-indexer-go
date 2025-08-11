package producers

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/nikitaNotFound/evm-indexer-go/pkg/workpool"
	"github.com/rs/zerolog/log"
)

const (
	BlocksTopicName = "blocks"
	RawTxsTopicName = "raw_txs"
)

type BlocksProducer struct {
	ethClient *ethclient.Client
}

func NewBlocksProducer(ethClient *ethclient.Client) *BlocksProducer {
	return &BlocksProducer{ethClient: ethClient}
}

type Block struct {
	Number    int64
	Hash      string
	Timestamp int64
}

type RawTx struct {
	Hash        string
	FromAddress string
	ToAddress   string
	Value       *big.Int
	Timestamp   int64
	BlockNumber int64
	InputData   string
	GasUsed     uint64
	MaxGasPrice *big.Int
}

// TODO: refactor to use shared workpool for blocks producer
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
	workPool := workpool.NewWorkPool[*types.Block](10)

	for i := trigger.StartBlock; i <= trigger.EndBlock; i++ {
		workPool.Enqueue(func() (*types.Block, error) {
			return p.loadBlockInfo(i, e.Ctx)
		})
	}

	workPool.WaitAndStop()

	for block := range workPool.Results() {
		if err := p.transformAndBroadcast(e, block); err != nil {
			l.Error().Err(err).Msg("failed to transform and broadcast block")
		}
	}

	return nil
}

func (p *BlocksProducer) handleSingleLoad(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().Str("component", "BlocksProducer").Str("method", "handleSingleLoad").Logger()

	blockInfo, err := p.loadBlockInfo(trigger.StartBlock, e.Ctx)
	if err != nil {
		l.Error().Err(err).Msg("failed to load block info from node")
		return err
	}

	if err := p.transformAndBroadcast(e, blockInfo); err != nil {
		l.Error().Err(err).Msg("failed to transform and broadcast block")
		return err
	}

	return nil
}

func (p *BlocksProducer) transformAndBroadcast(
	e engine.EngineCtx,
	blockInfo *types.Block,
) error {
	l := log.With().
		Str("component", "BlocksProducer").
		Str("method", "transformAndBroadcast").
		Int64("block_number", blockInfo.Number().Int64()).
		Logger()

	blockEvent := &Block{
		Number:    int64(blockInfo.Number().Int64()),
		Hash:      blockInfo.Hash().Hex(),
		Timestamp: blockInfo.ReceivedAt.UTC().Unix(),
	}

	txDataEvents := make([]models.ProducedDataEvent, len(blockInfo.Transactions()))
	for i, tx := range blockInfo.Transactions() {
		fromAddress, err := p.getSenderAddress(tx)
		if err != nil {
			l.Error().Err(err).
				Str("tx_hash", tx.Hash().Hex()).
				Msg("failed to recover sender address, continuing with empty address")
			return err
		}

		// Handle To address (can be nil for contract creation)
		toAddress := ""
		if tx.To() != nil {
			toAddress = tx.To().Hex()
		}

		txDataEvents[i] = models.NewProducedDataEvent(&RawTx{
			Hash:        tx.Hash().Hex(),
			FromAddress: fromAddress,
			ToAddress:   toAddress,
			Value:       tx.Value(),
			Timestamp:   blockInfo.ReceivedAt.UTC().Unix(),
			BlockNumber: int64(blockInfo.Number().Int64()),
			InputData:   hex.EncodeToString(tx.Data()),
			GasUsed:     tx.Gas(),
			MaxGasPrice: tx.GasPrice(),
		})
	}

	if err := e.BroadcastData(
		e.Ctx,
		BlocksTopicName,
		[]models.ProducedDataEvent{
			models.NewProducedDataEvent(blockEvent),
		},
	); err != nil {
		l.Error().Err(err).Msg("failed to broadcast data")
		return err
	}

	if err := e.BroadcastData(
		e.Ctx,
		RawTxsTopicName,
		txDataEvents,
	); err != nil {
		l.Error().Err(err).Msg("failed to broadcast data")
		return err
	}

	return nil
}

func (p *BlocksProducer) loadBlockInfo(
	blockNumber int64,
	ctx context.Context,
) (*types.Block, error) {
	l := log.With().Str("component", "BlocksProducer").Str("method", "loadBlockInfo").Logger()
	block, err := p.ethClient.BlockByNumber(ctx, big.NewInt(blockNumber))
	if err != nil {
		l.Error().Err(err).Msg("failed to load block info from node")
		return nil, err
	}

	return block, nil
}

// getSenderAddress recovers the sender address from a transaction
func (p *BlocksProducer) getSenderAddress(tx *types.Transaction) (string, error) {
	var signer types.Signer

	chainID := tx.ChainId()
	if chainID == nil || chainID.Sign() == 0 {
		// For legacy transactions without chain ID, use Homestead signer
		signer = types.HomesteadSigner{}
	} else {
		signer = types.LatestSignerForChainID(chainID)
	}

	sender, err := types.Sender(signer, tx)
	if err != nil {
		return "", err
	}

	return sender.Hex(), nil
}

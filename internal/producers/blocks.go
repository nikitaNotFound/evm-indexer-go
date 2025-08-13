package producers

import (
	"encoding/hex"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/nikitaNotFound/evm-indexer-go/pkg/smartlim"
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
	Number       int64
	Hash         string
	Timestamp    int64
	GasPrice     string
	BurntFees    string
	TotalFees    string
	TotalGas     uint64
	FeeRecipient string
}

type RawTx struct {
	Hash           string
	FromAddress    string
	ToAddress      string
	Value          *big.Int
	Timestamp      int64
	BlockNumber    int64
	InputData      string
	GasUsed        uint64
	GasPrice       *big.Int
	GasLimit       uint64
	MaxPriorityFee *big.Int
	MaxFee         *big.Int
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

	wg := sync.WaitGroup{}
	for i := trigger.StartBlock; i <= trigger.EndBlock; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			blockInfo, err := p.loadBlockInfo(e, i)
			if err != nil {
				l.Error().Err(err).Msg("failed to load block info from node")
				return
			}

			if err := p.transformAndBroadcast(e, blockInfo); err != nil {
				l.Error().Err(err).Msg("failed to transform and broadcast block")
				return
			}
		}()
	}
	wg.Wait()

	return nil
}

func (p *BlocksProducer) handleSingleLoad(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().Str("component", "BlocksProducer").Str("method", "handleSingleLoad").Logger()

	blockInfo, err := p.loadBlockInfo(e, trigger.StartBlock)
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

	mu := sync.Mutex{}
	totalGas := uint64(0)
	totalFees := big.NewInt(0)
	burntFees := big.NewInt(0)
	txDataEvents := make([]models.ProducedDataEvent, len(blockInfo.Transactions()))

	blockTransactions := blockInfo.Transactions()
	wg := sync.WaitGroup{}
	wg.Add(blockTransactions.Len())
	for i, tx := range blockTransactions {
		go func() {
			defer wg.Done()
			txData, err := p.loadTxData(e, tx, blockInfo)
			if err != nil {
				l.Error().Err(err).Msg("failed to load tx data")
				return
			}

			txDataEvents[i] = models.NewProducedDataEvent(txData)

			mu.Lock()
			defer mu.Unlock()
			totalGas += txData.GasUsed
			burntFees.Add(
				burntFees,
				tx.GasPrice().Mul(blockInfo.BaseFee(), big.NewInt(int64(txData.GasUsed))),
			)
			totalFees.Add(totalFees, tx.Cost())
		}()
	}
	wg.Wait()

	blockEvent := &Block{
		Number:       int64(blockInfo.Number().Int64()),
		Hash:         blockInfo.Hash().Hex(),
		Timestamp:    blockInfo.ReceivedAt.UTC().Unix(),
		GasPrice:     blockInfo.BaseFee().String(),
		TotalFees:    totalFees.String(),
		TotalGas:     totalGas,
		BurntFees:    burntFees.String(),
		FeeRecipient: blockInfo.Coinbase().Hex(),
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

func (p *BlocksProducer) loadTxData(
	e engine.EngineCtx,
	tx *types.Transaction,
	blockInfo *types.Block,
) (*RawTx, error) {
	l := log.With().
		Str("component", "BlocksProducer").
		Str("method", "loadTxData").
		Str("tx_hash", tx.Hash().Hex()).
		Logger()

	fromAddress, err := p.getSenderAddress(tx)
	if err != nil {
		l.Error().Err(err).
			Str("tx_hash", tx.Hash().Hex()).
			Msg("failed to recover sender address, continuing with empty address")
		return nil, err
	}

	toAddress := ""
	if tx.To() != nil {
		toAddress = tx.To().Hex()
	}

	receipt, err := smartlim.Process(e.Limiter, e.Ctx, func() (*types.Receipt, error) {
		return p.ethClient.TransactionReceipt(e.Ctx, tx.Hash())
	})
	if err != nil {
		l.Warn().Err(err).
			Str("tx_hash", tx.Hash().Hex()).
			Msg("failed to get transaction receipt")
		return nil, err
	}

	gasUsedForTx := receipt.GasUsed

	txDataEvent := &RawTx{
		Hash:           tx.Hash().Hex(),
		FromAddress:    fromAddress,
		ToAddress:      toAddress,
		Value:          tx.Value(),
		Timestamp:      blockInfo.ReceivedAt.UTC().Unix(),
		BlockNumber:    int64(blockInfo.Number().Int64()),
		InputData:      hex.EncodeToString(tx.Data()),
		GasUsed:        gasUsedForTx,
		GasPrice:       tx.GasPrice(),
		GasLimit:       tx.Gas(),
		MaxPriorityFee: tx.GasTipCap(),
		MaxFee:         tx.GasFeeCap(),
	}

	return txDataEvent, nil
}

func (p *BlocksProducer) loadBlockInfo(
	e engine.EngineCtx,
	blockNumber int64,
) (*types.Block, error) {
	l := log.With().Str("component", "BlocksProducer").Str("method", "loadBlockInfo").Logger()
	block, err := smartlim.Process(e.Limiter, e.Ctx, func() (*types.Block, error) {
		return p.ethClient.BlockByNumber(e.Ctx, big.NewInt(blockNumber))
	})
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

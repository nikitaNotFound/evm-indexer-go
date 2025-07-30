package indexers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/nikitaNotFound/evm-indexer-go/internal/producers"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storages/postgres"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storages/postgres/sqlcgen"
	"github.com/uptrace/bun/driver/pgdriver"
)

type RawTxsIndexer struct {
	pgStorage *postgres.Storage
}

func NewRawTxsIndexer(pgStorage *postgres.Storage) *RawTxsIndexer {
	return &RawTxsIndexer{pgStorage: pgStorage}
}

func (i *RawTxsIndexer) OnDataEvent(
	ctx context.Context,
	topic string,
	data models.ProducedDataEvent,
) error {
	rawTx, ok := data.Data.(*producers.RawTx)
	if !ok {
		return fmt.Errorf("invalid data type: %T", data)
	}

	if err := i.pgStorage.Queries.AddRawTx(ctx, sqlcgen.AddRawTxParams{
		Hash:        rawTx.Hash,
		FromAddress: rawTx.FromAddress,
		ToAddress:   rawTx.ToAddress,
		Value:       rawTx.Value.String(),
		Timestamp:   rawTx.Timestamp,
		BlockNumber: rawTx.BlockNumber,
		InputData:   rawTx.InputData,
		GasUsed:     strconv.FormatUint(rawTx.GasUsed, 10),
		MaxGasPrice: rawTx.MaxGasPrice.String(),
	}); err != nil {
		if pgerr, ok := err.(pgdriver.Error); ok && pgerr.IntegrityViolation() {
			return nil
		}
		return fmt.Errorf("failed to add raw tx: %w", err)
	}

	return nil
}

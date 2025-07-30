package indexers

import (
	"context"
	"fmt"

	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/nikitaNotFound/evm-indexer-go/internal/producers"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storages/postgres"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storages/postgres/sqlcgen"
	"github.com/uptrace/bun/driver/pgdriver"
)

type BlocksIndexer struct {
	pgStorage *postgres.Storage
}

func NewBlocksIndexer(pgStorage *postgres.Storage) *BlocksIndexer {
	return &BlocksIndexer{pgStorage: pgStorage}
}

func (i *BlocksIndexer) OnDataEvent(
	ctx context.Context,
	topic string,
	data models.ProducedDataEvent,
) error {
	blockInfo, ok := data.Data.(*producers.Block)
	if !ok {
		return fmt.Errorf("invalid data type: %T", data)
	}

	if err := i.pgStorage.Queries.AddBlock(ctx, sqlcgen.AddBlockParams{
		Number:    blockInfo.Number,
		Hash:      blockInfo.Hash,
		Timestamp: blockInfo.Timestamp,
	}); err != nil {
		if pgerr, ok := err.(pgdriver.Error); ok && pgerr.IntegrityViolation() {
			return nil
		}
		return fmt.Errorf("failed to add block: %w", err)
	}

	return nil
}

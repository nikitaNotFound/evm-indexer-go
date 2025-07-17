package indexers

import (
	"context"
	"fmt"

	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/nikitaNotFound/evm-indexer-go/internal/producers"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storage/postgres"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storage/postgres/sqlcgen"
	"github.com/rs/zerolog/log"
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
	l := log.With().Str("component", "BlocksIndexer").Str("method", "OnDataEvent").Logger()

	l.Info().Str("topic", topic).Interface("data", data).Msg("blocks indexer received data")

	blockInfo, ok := data.Data.(*producers.Block)
	if !ok {
		return fmt.Errorf("invalid data type: %T", data)
	}

	if err := i.pgStorage.Queries.AddBlock(ctx, sqlcgen.AddBlockParams{
		Number:    blockInfo.Number,
		Hash:      blockInfo.Hash,
		Timestamp: 0,
	}); err != nil {
		return fmt.Errorf("failed to add block: %w", err)
	}

	return nil
}

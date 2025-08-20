package indexers

import (
	"context"
	"fmt"

	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/nikitaNotFound/evm-indexer-go/internal/producers"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storages/postgres"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storages/postgres/sqlcgen"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun/driver/pgdriver"
)

type PancakeV2Indexer struct {
	pgStorage *postgres.Storage
}

func NewPancakeV2Indexer(pgStorage *postgres.Storage) *PancakeV2Indexer {
	return &PancakeV2Indexer{pgStorage: pgStorage}
}

func (i *PancakeV2Indexer) OnDataEvent(
	ctx context.Context,
	topic string,
	data models.ProducedDataEvent,
) error {
	l := log.With().Str("component", "PancakeV2Indexer").Str("method", "OnDataEvent").Logger()

	pancakeV2Pool, ok := data.Data.(*producers.PancakeV2Pool)
	if !ok {
		return fmt.Errorf("invalid data type: %T", data.Data)
	}

	if err := i.pgStorage.Queries.AddPancakeV2Pool(ctx, sqlcgen.AddPancakeV2PoolParams{
		Address: pancakeV2Pool.Address,
		Token0:  pancakeV2Pool.Token0,
		Token1:  pancakeV2Pool.Token1,
	}); err != nil {
		if pgerr, ok := err.(pgdriver.Error); ok && pgerr.IntegrityViolation() {
			l.Warn().Str("address", pancakeV2Pool.Address).Msg("pancake v2 pool already exists")
			return nil
		}
		return fmt.Errorf("failed to add pancake v2 pool: %w", err)
	}

	return nil
}

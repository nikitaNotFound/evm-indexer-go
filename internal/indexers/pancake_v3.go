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

type PancakeV3Indexer struct {
	pgStorage *postgres.Storage
}

func NewPancakeV3Indexer(pgStorage *postgres.Storage) *PancakeV3Indexer {
	return &PancakeV3Indexer{pgStorage: pgStorage}
}

func (i *PancakeV3Indexer) OnDataEvent(
	ctx context.Context,
	topic string,
	data models.ProducedDataEvent,
) error {
	l := log.With().Str("component", "PancakeV3Indexer").Str("method", "OnDataEvent").Logger()

	pancakeV3Pool, ok := data.Data.(*producers.PancakeV3Pool)
	if !ok {
		return fmt.Errorf("invalid data type: %T", data.Data)
	}

	if err := i.pgStorage.Queries.AddPancakeV3Pool(ctx, sqlcgen.AddPancakeV3PoolParams{
		Address:     pancakeV3Pool.Address,
		Token0:      pancakeV3Pool.Token0,
		Token1:      pancakeV3Pool.Token1,
		Fee:         pancakeV3Pool.Fee.String(),
		TickSpacing: pancakeV3Pool.TickSpacing.String(),
	}); err != nil {
		if pgerr, ok := err.(pgdriver.Error); ok && pgerr.IntegrityViolation() {
			l.Warn().Str("address", pancakeV3Pool.Address).Msg("pancake v3 pool already exists")
			return nil
		}
		return fmt.Errorf("failed to add pancake v3 pool: %w", err)
	}

	return nil
}
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

type SushiSwapV2Indexer struct {
	pgStorage *postgres.Storage
}

func NewSushiSwapV2Indexer(pgStorage *postgres.Storage) *SushiSwapV2Indexer {
	return &SushiSwapV2Indexer{pgStorage: pgStorage}
}

func (i *SushiSwapV2Indexer) OnDataEvent(
	ctx context.Context,
	topic string,
	data models.ProducedDataEvent,
) error {
	l := log.With().Str("component", "SushiSwapV2Indexer").Str("method", "OnDataEvent").Logger()

	sushiSwapV2Pool, ok := data.Data.(*producers.SushiSwapV2Pool)
	if !ok {
		return fmt.Errorf("invalid data type: %T", data.Data)
	}

	if err := i.pgStorage.Queries.AddSushiSwapV2Pool(ctx, sqlcgen.AddSushiSwapV2PoolParams{
		Address: sushiSwapV2Pool.Address,
		Token0:  sushiSwapV2Pool.Token0,
		Token1:  sushiSwapV2Pool.Token1,
	}); err != nil {
		if pgerr, ok := err.(pgdriver.Error); ok && pgerr.IntegrityViolation() {
			l.Warn().Str("address", sushiSwapV2Pool.Address).Msg("sushiswap v2 pool already exists")
			return nil
		}
		return fmt.Errorf("failed to add sushiswap v2 pool: %w", err)
	}

	return nil
}
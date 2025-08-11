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

type UniswapV3Indexer struct {
	pgStorage *postgres.Storage
}

func NewUniswapV3Indexer(pgStorage *postgres.Storage) *UniswapV3Indexer {
	return &UniswapV3Indexer{pgStorage: pgStorage}
}

func (i *UniswapV3Indexer) OnDataEvent(
	ctx context.Context,
	topic string,
	data models.ProducedDataEvent,
) error {
	l := log.With().Str("component", "UniswapV2Indexer").Str("method", "OnDataEvent").Logger()

	uniswapV3Pool, ok := data.Data.(*producers.UniswapV3Pool)
	if !ok {
		return fmt.Errorf("invalid data type: %T", data.Data)
	}

	if err := i.pgStorage.Queries.AddUniswapV3Pool(ctx, sqlcgen.AddUniswapV3PoolParams{
		Address:     uniswapV3Pool.Address,
		Token0:      uniswapV3Pool.Token0,
		Token1:      uniswapV3Pool.Token1,
		Fee:         uniswapV3Pool.Fee.String(),
		TickSpacing: uniswapV3Pool.TickSpacing.String(),
	}); err != nil {
		if pgerr, ok := err.(pgdriver.Error); ok && pgerr.IntegrityViolation() {
			l.Warn().Str("address", uniswapV3Pool.Address).Msg("uniswap v3 pool already exists")
			return nil
		}
		return fmt.Errorf("failed to add uniswap v3 pool: %w", err)
	}

	return nil
}

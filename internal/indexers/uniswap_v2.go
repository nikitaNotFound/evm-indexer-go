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

type UniswapV2Indexer struct {
	pgStorage *postgres.Storage
}

func NewUniswapV2Indexer(pgStorage *postgres.Storage) *UniswapV2Indexer {
	return &UniswapV2Indexer{pgStorage: pgStorage}
}

func (i *UniswapV2Indexer) OnDataEvent(
	ctx context.Context,
	topic string,
	data models.ProducedDataEvent,
) error {
	l := log.With().Str("component", "UniswapV2Indexer").Str("method", "OnDataEvent").Logger()

	uniswapV2Pool, ok := data.Data.(*producers.UniswapV2Pool)
	if !ok {
		return fmt.Errorf("invalid data type: %T", data.Data)
	}

	if err := i.pgStorage.Queries.AddUniswapV2Pool(ctx, sqlcgen.AddUniswapV2PoolParams{
		Address: uniswapV2Pool.Address,
		Token0:  uniswapV2Pool.Token0,
		Token1:  uniswapV2Pool.Token1,
	}); err != nil {
		if pgerr, ok := err.(pgdriver.Error); ok && pgerr.IntegrityViolation() {
			l.Warn().Str("address", uniswapV2Pool.Address).Msg("uniswap v2 pool already exists")
			return nil
		}
		return fmt.Errorf("failed to add uniswap v2 pool: %w", err)
	}

	return nil
}

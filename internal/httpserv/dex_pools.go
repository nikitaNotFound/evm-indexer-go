package httpserv

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikitaNotFound/evm-indexer-go/internal/apigen"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storages/postgres/sqlcgen"
	"github.com/rs/zerolog/log"
)

// GetAllPools retrieves all indexed pools with pagination
func (s *HTTPServer) GetAllPools(ctx echo.Context, params apigen.GetAllPoolsParams) error {
	l := log.With().Str("service", "HttpServer").Str("method", "GetAllPools").Logger()

	if params.Limit != nil {
		l = l.With().Int("limit", *params.Limit).Logger()
	}
	if params.Offset != nil {
		l = l.With().Int("offset", *params.Offset).Logger()
	}

	l.Debug().Msg("retrieving pools data")

	// Validate parameters
	limit := 100 // default
	if params.Limit != nil {
		if *params.Limit < 1 || *params.Limit > 1000 {
			l.Warn().Msg("invalid limit parameter")
			return ctx.JSON(http.StatusBadRequest, apigen.Error{
				Error:   "Invalid limit parameter",
				Code:    "INVALID_LIMIT",
				Details: stringPtr("Limit must be between 1 and 1000"),
			})
		}
		limit = *params.Limit
	}

	offset := 0 // default
	if params.Offset != nil {
		if *params.Offset < 0 {
			l.Warn().Msg("invalid offset parameter")
			return ctx.JSON(http.StatusBadRequest, apigen.Error{
				Error:   "Invalid offset parameter",
				Code:    "INVALID_OFFSET",
				Details: stringPtr("Offset must be non-negative"),
			})
		}
		offset = *params.Offset
	}

	ctx_req := ctx.Request().Context()

	// Get all pools (both V2 and V3)
	pools, total, err := s.getAllPoolsCombined(ctx_req, int32(limit), int32(offset))
	if err != nil {
		l.Error().Err(err).Msg("failed to get pools from database")
		return ctx.JSON(http.StatusInternalServerError, apigen.Error{
			Error:   "Internal server error",
			Code:    "DATABASE_ERROR",
			Details: stringPtr("Failed to retrieve pools"),
		})
	}

	response := apigen.PoolsResponse{
		Pools:  pools,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}

	l.Debug().Int("total_pools", total).Int("returned_pools", len(pools)).
		Msg("pools data retrieved successfully")
	return ctx.JSON(http.StatusOK, response)
}

// getAllPoolsCombined retrieves both V2 and V3 pools with combined pagination
func (s *HTTPServer) getAllPoolsCombined(
	ctx_req context.Context, limit, offset int32,
) ([]apigen.Pool, int, error) {
	// Get total counts first
	v2Count, err := s.storage.Queries.CountUniswapV2Pools(ctx_req)
	if err != nil {
		return nil, 0, err
	}

	v3Count, err := s.storage.Queries.CountUniswapV3Pools(ctx_req)
	if err != nil {
		return nil, 0, err
	}

	totalCount := int(v2Count + v3Count)

	// Simple strategy: get V2 pools first, then V3 pools
	var allPools []apigen.Pool

	// Calculate how many V2 pools to fetch
	if offset < int32(v2Count) {
		v2Limit := limit
		if int64(offset+limit) > v2Count {
			v2Limit = int32(v2Count) - offset
		}

		v2Pools, err := s.storage.Queries.GetAllUniswapV2Pools(
			ctx_req, sqlcgen.GetAllUniswapV2PoolsParams{
				Limit:  v2Limit,
				Offset: offset,
			})
		if err != nil {
			return nil, 0, err
		}

		for _, dbPool := range v2Pools {
			pool := apigen.Pool{}
			if err := pool.FromUniswapV2Pool(apigen.UniswapV2Pool{
				Address: dbPool.Address,
				Token0:  dbPool.Token0,
				Token1:  dbPool.Token1,
				Version: "v2",
			}); err != nil {
				return nil, 0, err
			}
			allPools = append(allPools, pool)
		}
	}

	// Calculate how many V3 pools to fetch
	if len(allPools) < int(limit) && offset+int32(len(allPools)) < int32(totalCount) {
		remainingLimit := limit - int32(len(allPools))
		v3Offset := int32(0)
		if offset > int32(v2Count) {
			v3Offset = offset - int32(v2Count)
		}

		v3Pools, err := s.storage.Queries.GetAllUniswapV3Pools(
			ctx_req, sqlcgen.GetAllUniswapV3PoolsParams{
				Limit:  remainingLimit,
				Offset: v3Offset,
			})
		if err != nil {
			return nil, 0, err
		}

		for _, dbPool := range v3Pools {
			pool := apigen.Pool{}
			if err := pool.FromUniswapV3Pool(apigen.UniswapV3Pool{
				Address:     dbPool.Address,
				Token0:      dbPool.Token0,
				Token1:      dbPool.Token1,
				Fee:         dbPool.Fee,
				TickSpacing: dbPool.TickSpacing,
				Version:     "v3",
			}); err != nil {
				return nil, 0, err
			}
			allPools = append(allPools, pool)
		}
	}

	return allPools, totalCount, nil
}

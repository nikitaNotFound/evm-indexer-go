package httpserv

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikitaNotFound/evm-indexer-go/internal/apigen"
	"github.com/rs/zerolog/log"
)

// GetBlockByNumber retrieves block data by block number
func (s *HTTPServer) GetBlockByNumber(e echo.Context, number int64) error {
	ctx := e.Request().Context()

	l := log.With().Str("service", "HttpServer").Str("method", "GetBlockByNumber").
		Int64("block_number", number).Logger()

	if number < 0 {
		return e.JSON(http.StatusBadRequest, apigen.Error{
			Error:   "Invalid block number",
			Code:    "INVALID_BLOCK_NUMBER",
			Details: stringPtr("Block number must be non-negative"),
		})
	}

	dbBlock, err := s.storage.Queries.GetBlockByNumber(ctx, number)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.JSON(http.StatusNotFound, apigen.Error{
				Error:   "Block not found",
				Code:    "BLOCK_NOT_FOUND",
				Details: stringPtr("Block with the specified number does not exist"),
			})
		}

		l.Error().Err(err).Msg("failed to get block from database")
		return e.JSON(http.StatusInternalServerError, apigen.Error{
			Error:   "Internal server error",
			Code:    "DATABASE_ERROR",
			Details: stringPtr("Failed to retrieve block data"),
		})
	}

	block := apigen.Block{
		Number:       dbBlock.Number,
		Hash:         dbBlock.Hash,
		Timestamp:    dbBlock.Timestamp,
		GasPrice:     dbBlock.GasPrice,
		BurntFees:    dbBlock.BurntFees,
		TotalFees:    dbBlock.TotalFees,
		TotalGas:     uint64(dbBlock.TotalGas),
		FeeRecipient: dbBlock.FeeRecipient,
	}

	return e.JSON(http.StatusOK, block)
}

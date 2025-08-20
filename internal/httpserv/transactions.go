package httpserv

import (
	"database/sql"
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nikitaNotFound/evm-indexer-go/internal/apigen"
	"github.com/rs/zerolog/log"
)

var txHashPattern = regexp.MustCompile(`^0x[a-fA-F0-9]{64}$`)

// GetTransactionByHash retrieves transaction data by transaction hash
func (s *HTTPServer) GetTransactionByHash(ctx echo.Context, hash string) error {
	l := log.With().Str("service", "HttpServer").Str("method", "GetTransactionByHash").
		Str("tx_hash", hash).Logger()

	if !txHashPattern.MatchString(hash) {
		return ctx.JSON(http.StatusBadRequest, apigen.Error{
			Error:   "Invalid transaction hash format",
			Code:    "INVALID_TX_HASH",
			Details: stringPtr("Transaction hash must be 64-character hex string with 0x prefix"),
		})
	}

	dbTx, err := s.storage.Queries.GetTransactionByHash(ctx.Request().Context(), hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, apigen.Error{
				Error:   "Transaction not found",
				Code:    "TRANSACTION_NOT_FOUND",
				Details: stringPtr("Transaction with the specified hash does not exist"),
			})
		}

		l.Error().Err(err).Msg("failed to get transaction from database")
		return ctx.JSON(http.StatusInternalServerError, apigen.Error{
			Error:   "Internal server error",
			Code:    "DATABASE_ERROR",
			Details: stringPtr("Failed to retrieve transaction data"),
		})
	}

	transaction := apigen.RawTx{
		Hash:           dbTx.Hash,
		FromAddress:    dbTx.FromAddress,
		ToAddress:      dbTx.ToAddress,
		Value:          dbTx.Value,
		Timestamp:      dbTx.Timestamp,
		BlockNumber:    dbTx.BlockNumber,
		InputData:      dbTx.InputData,
		GasUsed:        parseUint64(dbTx.GasUsed),
		GasPrice:       dbTx.GasPrice,
		GasLimit:       parseUint64(dbTx.GasLimit),
		MaxPriorityFee: dbTx.MaxPriorityFee,
		MaxFee:         dbTx.MaxFee,
	}

	return ctx.JSON(http.StatusOK, transaction)
}

func parseUint64(s string) uint64 {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

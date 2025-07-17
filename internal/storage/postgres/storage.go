package postgres

import (
	"context"
	"database/sql"

	"github.com/nikitaNotFound/evm-indexer-go/internal/storage/postgres/sqlcgen"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Storage struct {
	Queries *sqlcgen.Queries
	db      *sql.DB
}

func NewStorage(dsn string) *Storage {
	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	return &Storage{
		Queries: sqlcgen.New(db),
		db:      db,
	}
}

func (s *Storage) Close() error {
	return s.db.Close()
}

// GetDB returns the database connection for transactions
func (s *Storage) GetDB() *sql.DB {
	return s.db
}

// WithTx returns queries with transaction
func (s *Storage) WithTx(tx *sql.Tx) *sqlcgen.Queries {
	return s.Queries.WithTx(tx)
}

func (s *Storage) BeginTx(ctx context.Context) (*sql.Tx, *sqlcgen.Queries, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	qtx := s.Queries.WithTx(tx)
	return tx, qtx, nil
}

func (s *Storage) Ping() error {
	return s.db.Ping()
}

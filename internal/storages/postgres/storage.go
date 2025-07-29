package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/nikitaNotFound/evm-indexer-go/internal/storages/postgres/sqlcgen"
	"github.com/uptrace/bun/driver/pgdriver"
)

type options struct {
	createDBIfNotExists bool
	maxOpenConns        int
	maxIdleConns        int
	connMaxLifetime     time.Duration
	connMaxIdleTime     time.Duration
}

func defaultOptions() *options {
	return &options{
		createDBIfNotExists: false,
		maxOpenConns:        25,              // Max 25 open connections
		maxIdleConns:        5,               // Keep 5 idle connections
		connMaxLifetime:     5 * time.Minute, // Connections live for 5 minutes max
		connMaxIdleTime:     1 * time.Minute, // Idle connections timeout after 1 minute
	}
}

type Option func(*options)

func WithCreateDBIfNotExists() Option {
	return func(o *options) {
		o.createDBIfNotExists = true
	}
}

// WithMaxOpenConns sets the maximum number of open connections
func WithMaxOpenConns(maxOpen int) Option {
	return func(o *options) {
		o.maxOpenConns = maxOpen
	}
}

// WithMaxIdleConns sets the maximum number of idle connections
func WithMaxIdleConns(maxIdle int) Option {
	return func(o *options) {
		o.maxIdleConns = maxIdle
	}
}

type Storage struct {
	Queries *sqlcgen.Queries
	db      *sql.DB
	dsn     string
}

func NewStorage(dsn string, opts ...Option) (*Storage, error) {
	opt := defaultOptions()
	for _, o := range opts {
		o(opt)
	}

	if opt.createDBIfNotExists {
		if err := createDatabaseIfNotExists(dsn); err != nil {
			return nil, fmt.Errorf("failed to create database: %w", err)
		}
	}

	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db.SetMaxOpenConns(opt.maxOpenConns)
	db.SetMaxIdleConns(opt.maxIdleConns)
	db.SetConnMaxLifetime(opt.connMaxLifetime)
	db.SetConnMaxIdleTime(opt.connMaxIdleTime)

	return &Storage{
		Queries: sqlcgen.New(db),
		db:      db,
		dsn:     dsn,
	}, nil
}

// PoolStats returns connection pool statistics
func (s *Storage) PoolStats() sql.DBStats {
	return s.db.Stats()
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

func (s *Storage) Migrate() error {
	return migrateDb(s.dsn)
}

// createDatabaseIfNotExists parses the DSN, connects to PostgreSQL without specifying
// the target database, checks if the database exists, and creates it if it doesn't
func createDatabaseIfNotExists(dsn string) error {
	dbName, adminDSN, err := parseDSNForDBCreation(dsn)
	if err != nil {
		return fmt.Errorf("failed to parse DSN: %w", err)
	}

	adminDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(adminDSN)))
	defer adminDB.Close()

	exists, err := databaseExists(adminDB, dbName)
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	if !exists {
		query := fmt.Sprintf("CREATE DATABASE %s", dbName)
		_, err := adminDB.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to create database '%s': %w", dbName, err)
		}
	}

	return nil
}

// parseDSNForDBCreation extracts the database name and creates an admin DSN
// that connects to the 'postgres' database instead
func parseDSNForDBCreation(dsn string) (dbName, adminDSN string, err error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return "", "", err
	}

	dbName = strings.TrimPrefix(u.Path, "/")
	if dbName == "" {
		return "", "", fmt.Errorf("no database name found in DSN")
	}

	adminURL := *u
	adminURL.Path = "/postgres"
	adminDSN = adminURL.String()

	return dbName, adminDSN, nil
}

// databaseExists checks if a database with the given name exists
func databaseExists(db *sql.DB, dbName string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)"
	err := db.QueryRow(query, dbName).Scan(&exists)
	return exists, err
}

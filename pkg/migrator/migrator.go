package migrator

import (
	"database/sql"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
)

type Migrator struct {
	db         *sql.DB
	migrations fs.FS
	tableName  string
	driverType string
}

type Options struct {
	TableName  string
	DriverType string
}

// NewMigrator creates a new migrator instance
func NewMigrator(db *sql.DB, migrationsFS fs.FS, opts *Options) *Migrator {
	if opts == nil {
		opts = &Options{
			TableName:  "schema_migrations",
			DriverType: "postgres",
		}
	}

	if opts.TableName == "" {
		opts.TableName = "schema_migrations"
	}
	if opts.DriverType == "" {
		opts.DriverType = "postgres"
	}

	return &Migrator{
		db:         db,
		migrations: migrationsFS,
		tableName:  opts.TableName,
		driverType: opts.DriverType,
	}
}

// Migrate runs all pending migrations using flyway notation (V1__description.sql)
func (m *Migrator) Migrate() error {
	l := log.With().Str("service", "migrator").Str("method", "Migrate").Logger()

	// Create migration tracking table
	if err := m.createMigrationTable(); err != nil {
		return fmt.Errorf("failed to create migration table: %w", err)
	}

	// Get list of migration files
	migrationFiles, err := m.getMigrationFiles()
	if err != nil {
		return fmt.Errorf("failed to get migration files: %w", err)
	}

	// Execute migrations
	for _, filename := range migrationFiles {
		if applied, err := m.isMigrationApplied(filename); err != nil {
			return fmt.Errorf("failed to check migration status for %s: %w", filename, err)
		} else if applied {
			l.Debug().Str("migration", filename).Msg("Migration already applied, skipping")
			continue
		}

		l.Info().Str("migration", filename).Msg("Applying migration")
		if err := m.executeMigration(filename); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", filename, err)
		}

		if err := m.markMigrationApplied(filename); err != nil {
			return fmt.Errorf("failed to mark migration as applied %s: %w", filename, err)
		}
		l.Info().Str("migration", filename).Msg("Migration applied successfully")
	}

	return nil
}

// createMigrationTable creates the migration tracking table
func (m *Migrator) createMigrationTable() error {
	// Check if table exists and has correct schema
	var exists bool
	err := m.db.QueryRow(`
		SELECT EXISTS (
			SELECT FROM information_schema.tables 
			WHERE table_name = $1
		)`, m.tableName).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		// Check if schema is correct (VARCHAR version column)
		var columnType string
		err = m.db.QueryRow(`
			SELECT data_type 
			FROM information_schema.columns 
			WHERE table_name = $1 AND column_name = 'version'
		`, m.tableName).Scan(&columnType)

		// If version column is not character varying, drop and recreate
		if err != nil || columnType != "character varying" {
			_, err = m.db.Exec(fmt.Sprintf("DROP TABLE %s", m.tableName))
			if err != nil {
				return err
			}
			exists = false
		}
	}

	// Create table if it doesn't exist or was dropped
	if !exists {
		query := fmt.Sprintf(`
			CREATE TABLE %s (
				version VARCHAR(255) PRIMARY KEY,
				applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			)`, m.tableName)
		_, err = m.db.Exec(query)
		return err
	}

	return nil
}

// getMigrationFiles returns sorted list of migration files
func (m *Migrator) getMigrationFiles() ([]string, error) {
	var files []string
	err := fs.WalkDir(m.migrations, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".sql") {
			files = append(files, filepath.Base(path))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Sort files to ensure consistent execution order
	sort.Strings(files)
	return files, nil
}

// isMigrationApplied checks if migration has already been applied
func (m *Migrator) isMigrationApplied(filename string) (bool, error) {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE version = $1", m.tableName)
	err := m.db.QueryRow(query, filename).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// executeMigration executes a migration file
func (m *Migrator) executeMigration(filename string) error {
	content, err := fs.ReadFile(m.migrations, filename)
	if err != nil {
		return err
	}

	_, err = m.db.Exec(string(content))
	return err
}

// markMigrationApplied marks migration as applied in tracking table
func (m *Migrator) markMigrationApplied(filename string) error {
	query := fmt.Sprintf("INSERT INTO %s (version) VALUES ($1)", m.tableName)
	_, err := m.db.Exec(query, filename)
	return err
}

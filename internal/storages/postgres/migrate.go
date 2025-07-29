package postgres

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"

	_ "github.com/lib/pq"
	"github.com/nikitaNotFound/evm-indexer-go/pkg/migrator"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// migrateDb runs database migrations using flyway notation (V1__description.sql)
func migrateDb(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	migrations, err := fs.Sub(embedMigrations, "migrations")
	if err != nil {
		return fmt.Errorf("failed to get migrations filesystem: %w", err)
	}

	m := migrator.NewMigrator(db, migrations, &migrator.Options{
		TableName:  "schema_migrations",
		DriverType: "postgres",
	})

	return m.Migrate()
}

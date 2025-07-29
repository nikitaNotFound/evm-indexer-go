# Migrator Package

A simple, lightweight database migration system with support for flyway notation.

## Features

- ✅ **Flyway notation**: `V1__description.sql`, `V2__description.sql`
- ✅ **Up-only migrations**: Perfect for indexers and append-only systems
- ✅ **Embedded filesystem support**: Migrations built into binary
- ✅ **Database agnostic**: Works with any SQL database
- ✅ **Migration tracking**: Automatic duplicate prevention
- ✅ **Sorted execution**: Migrations run in predictable order

## Usage

```go
package main

import (
    "database/sql"
    "embed"
    "io/fs"
    
    _ "github.com/lib/pq"
    "github.com/yourproject/pkg/migrator"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
    // Connect to database
    db, err := sql.Open("postgres", "postgres://localhost/mydb")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Get migrations filesystem
    migrations, err := fs.Sub(embedMigrations, "migrations")
    if err != nil {
        panic(err)
    }

    // Create and run migrator
    m := migrator.NewMigrator(db, migrations, &migrator.Options{
        TableName:  "schema_migrations",
        DriverType: "postgres",
    })

    if err := m.Migrate(); err != nil {
        panic(err)
    }
}
```

## Migration File Format

Create migration files following flyway notation:

```
migrations/
├── V1__create_users_table.sql
├── V2__add_indexes.sql
├── V3__create_products_table.sql
└── V4__alter_users_add_email.sql
```

### Example Migration Files

**V1__create_users_table.sql:**
```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**V2__add_indexes.sql:**
```sql
CREATE INDEX idx_users_name ON users(name);
CREATE INDEX idx_users_created_at ON users(created_at);
```

## Options

```go
type Options struct {
    TableName  string  // Migration tracking table name (default: "schema_migrations")
    DriverType string  // Database driver type (default: "postgres")
}
```

## Migration Tracking

The migrator automatically creates a tracking table to prevent duplicate migrations:

```sql
CREATE TABLE schema_migrations (
    version VARCHAR(255) PRIMARY KEY,
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Error Handling

All errors are wrapped with context for easy debugging:

```go
if err := m.Migrate(); err != nil {
    log.Fatal().Err(err).Msg("Migration failed")
}
```

Typical error scenarios:
- Database connection issues
- SQL syntax errors in migration files
- Missing migration files
- Filesystem access problems 
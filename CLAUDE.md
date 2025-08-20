# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Essential Commands

```bash
# Run the indexer
make run

# Build the binary
make build

# Run tests
make test

# Run linter
make lint

# Fix linting issues automatically
make lint-fix

# Format code
make fmt

# Generate API code from OpenAPI spec
make gen-api

# Generate Go bindings from smart contract ABIs
make abigen

# Install development tools (golangci-lint, goimports)
make install-tools
```

## Architecture Overview

This is an EVM blockchain indexer that fetches data through Node RPCs and stores it in PostgreSQL. The system follows a producer-consumer pattern with topic-based event distribution.

### Core Flow
1. **Producers** fetch blockchain data from Ethereum nodes via RPC
2. **Engine** orchestrates producers and broadcasts events to topics
3. **IndexersGate** manages topic subscriptions and event distribution
4. **Indexers** consume events and persist data to PostgreSQL
5. **HTTP API** serves indexed data via REST endpoints (optional)

### Key Components

**Engine** (`internal/engine/`)
- Central orchestrator that manages data flow
- Controls producer execution with configurable block ranges
- Implements rate limiting via SmartLimiter to avoid provider 429 errors
- Broadcasts produced data to topic subscribers

**Producers** (`internal/producers/`)
- Fetch blockchain data from Ethereum nodes
- Current implementations: blocks, raw transactions, Uniswap V2/V3 pools
- Each producer implements the `DataProducer` interface with `OnProduce` method

**Indexers** (`internal/indexers/`)
- Subscribe to topics and persist data to storage
- Each indexer implements the `Indexer` interface with `OnDataEvent` method
- Handle batch inserts for efficiency

**Storage** (`internal/storages/postgres/`)
- PostgreSQL storage with automatic migrations
- Uses SQLC for type-safe SQL queries (see `sqlc.yaml` configuration)
- Query definitions in `queries/*.sql` generate Go code in `sqlcgen/`

**HTTP Server** (`internal/httpserv/`)
- Optional REST API for accessing indexed data
- Generated from OpenAPI spec (`spec/openapi.yaml`)
- Uses Echo framework with Swagger documentation

### Configuration

The indexer uses `config.json` with the following structure:
- `network`: RPC URL, block range, batch size, update frequency
- `pg_storage`: PostgreSQL connection string
- `api`: HTTP server settings (optional)

### Database Schema

Tables created via migrations:
- `blocks`: Block data with gas info and fees
- `raw_txs`: Transaction details with gas usage
- `uniswap_v2_pools`: V2 pool addresses and tokens
- `uniswap_v3_pools`: V3 pools with fee tiers

## Code Style Guidelines

Based on `.cursor/rules/code-rules.mdc`:

1. Handle and log all errors. Business layer services return custom errors identifiable by API layer
2. Use short, clear naming (e.g., `tx` instead of `transaction`)
3. Omit `Get` prefix in function names (e.g., `UserTransactions` not `GetUserTransactions`)
4. Add concise function comments describing purpose
5. Use index-based slice population instead of `append` when possible
6. Use nested zerolog instances for detailed context:
```go
l := log.With().Str("service", "user").Str("method", "SomeComplexLogic").Str("user_id", userID).Logger()
```
7. Log errors with `.Err(err)` method, not `%s` formatting

## Adding New Indexers

1. Create producer in `internal/producers/` implementing `DataProducer`
2. Create indexer in `internal/indexers/` implementing `Indexer`
3. Add SQL queries in `internal/storages/postgres/queries/`
4. Run `sqlc generate` to create type-safe query methods
5. Register producer and indexer in `internal/apprun/runner.go`
6. Create topic and subscribe indexer to it

## Smart Contract Integration

1. Place ABI files in `internal/smartcontracts/abi/`
2. Run `make abigen` to generate Go bindings
3. Generated code appears in `internal/smartcontracts/abigen/`
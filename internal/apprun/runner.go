package apprun

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"

	"github.com/nikitaNotFound/evm-indexer-go/internal/config"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine"
	"github.com/nikitaNotFound/evm-indexer-go/internal/indexers"
	"github.com/nikitaNotFound/evm-indexer-go/internal/producers"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storages/postgres"
)

// StartEVMIndexer starts the EVM indexer with graceful shutdown support
func StartEVMIndexer() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}

	setupLogger(cfg)
	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		log.Info().Str("signal", sig.String()).Msg("received shutdown signal, starting graceful shutdown")
		cancel()
	}()

	ethClient, err := ethclient.Dial(cfg.NetworkConfig.RpcUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to node")
	}

	pgStorage, err := postgres.NewStorage(
		cfg.PGStorage.ConnectionString,
		postgres.WithCreateDBIfNotExists(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create postgres storage")
	}

	if err := pgStorage.Migrate(); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate database")
	}

	blocksProducer := producers.NewBlocksProducer(ethClient)
	blocksIndexer := indexers.NewBlocksIndexer(pgStorage)

	engine := engine.CreateEngine(cfg, []engine.DataProducer{
		blocksProducer,
	})

	if err := engine.IndexersGate().CreateTopic("blocks"); err != nil {
		log.Fatal().Err(err).Msg("failed to create blocks topic")
	}

	if err := engine.IndexersGate().Subscribe("blocks", blocksIndexer); err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to blocks topic")
	}

	if err := engine.Start(ctx); err != nil {
		if err == context.Canceled {
			log.Info().Msg("indexer stopped gracefully")
		} else {
			log.Error().Err(err).Msg("indexer stopped with error")
		}
	}
}

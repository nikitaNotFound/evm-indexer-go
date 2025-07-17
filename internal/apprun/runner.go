package apprun

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"

	"github.com/nikitaNotFound/evm-indexer-go/internal/config"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine"
	"github.com/nikitaNotFound/evm-indexer-go/internal/indexers"
	"github.com/nikitaNotFound/evm-indexer-go/internal/producers"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storage/postgres"
)

func StartEVMIndexer() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	setupLogger(cfg)

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
		log.Fatal().Err(err).Msg("engine stopped")
	}
}

package apprun

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/nikitaNotFound/evm-indexer-go/internal/config"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine"
	"github.com/nikitaNotFound/evm-indexer-go/internal/indexers"
	"github.com/nikitaNotFound/evm-indexer-go/internal/producers"
)

func StartEVMIndexer() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	setupLogger(cfg)

	blocksProducer := producers.NewBlocksProducer()
	blocksIndexer := indexers.NewBlocksIndexer()

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

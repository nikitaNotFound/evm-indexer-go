package apprun

import (
	"context"

	"github.com/nikitaNotFound/evm-indexer-go/internal/core"
	"github.com/nikitaNotFound/evm-indexer-go/internal/core/config"
	"github.com/nikitaNotFound/evm-indexer-go/internal/core/contracts"
	"github.com/rs/zerolog/log"
)

func StartEVMIndexer() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}

	ctx, _ := context.WithCancel(context.Background())
	setupLogger(cfg)

	engine := core.CreateEngine(cfg, []contracts.DataProducer{})
	if err := engine.Start(ctx); err != nil {
		log.Fatal().Err(err).Msg("engine stopped")
	}
}

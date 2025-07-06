package apprun

import (
	"os"

	"github.com/nikitaNotFound/evm-indexer-go/internal/core/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func setupLogger(cfg *config.Config) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if cfg.IsDebug() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Info().Msg("╔═══════════════════════════════════╗")
		log.Info().Msg("║         DEBUG MODE ENABLED        ║")
		log.Info().Msg("╚═══════════════════════════════════╝")
		log.Info().Msg("Log Level: DEBUG")
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("Log Level: INFO")
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

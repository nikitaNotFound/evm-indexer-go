package core

import (
	"context"

	"github.com/nikitaNotFound/evm-indexer-go/internal/core/config"
	"github.com/nikitaNotFound/evm-indexer-go/internal/core/contracts"
	"github.com/nikitaNotFound/evm-indexer-go/internal/core/models"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type Engine struct {
	dataProducers []contracts.DataProducer
	cfg           *config.Config
}

func CreateEngine(cfg *config.Config, dp []contracts.DataProducer) Engine {
	return Engine{
		dataProducers: dp,
		cfg:           cfg,
	}
}

func (e *Engine) Start(ctx context.Context) error {
	l := log.With().Str("component", "Engine").Str("method", "Start").Logger()

	if err := e.initialSync(); err != nil {
		l.Error().Err(err).Msg("starting engine failed on initial sync")
	}

	return nil
}

func (e *Engine) initialSync() error {
	l := log.With().Str("component", "Engine").Str("method", "initialSync").Logger()
	errGroup := new(errgroup.Group)

	for _, p := range e.dataProducers {
		errGroup.Go(func() error {
			return p.OnProduceTrigger(models.TriggerDataProduce{
				StartBlock: e.cfg.NetworkConfig.StartBlock,
				EndBlock:   e.cfg.NetworkConfig.EndBlock,
			})
		})
	}

	if err := errGroup.Wait(); err != nil {
		l.Error().Err(err).Msg("some of data producers failed")
		return err
	}

	return nil
}

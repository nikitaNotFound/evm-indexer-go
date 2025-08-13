package engine

import (
	"context"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"github.com/nikitaNotFound/evm-indexer-go/internal/config"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/nikitaNotFound/evm-indexer-go/pkg/smartlim"
)

type Engine struct {
	dataProducers   []DataProducer
	cfg             *config.Config
	indexersGate    *IndexersGate
	providerLimiter *smartlim.SmartLimiter
}

type EngineCtx struct {
	Engine        *Engine
	Ctx           context.Context
	BroadcastData func(ctx context.Context, topic string, data []models.ProducedDataEvent) error
	Limiter       *smartlim.SmartLimiter
}

func CreateEngine(cfg *config.Config, dp []DataProducer) Engine {
	limiter := smartlim.StartSmartLimiter(cfg.NetworkConfig.Rps, 1)

	return Engine{
		dataProducers:   dp,
		cfg:             cfg,
		indexersGate:    NewIndexersGate(),
		providerLimiter: limiter,
	}
}

func (e *Engine) IndexersGate() *IndexersGate {
	return e.indexersGate
}

func (e *Engine) Start(ctx context.Context) error {
	l := log.With().Str("component", "Engine").Str("method", "Start").Logger()

	if err := e.initialSync(ctx); err != nil {
		l.Error().Err(err).Msg("starting engine failed on initial sync")
	}

	e.indexersGate.WaitFinish()

	return nil
}

func (e *Engine) Stop() {
	e.indexersGate.Stop()
}

func (e *Engine) initialSync(ctx context.Context) error {
	l := log.With().Str("component", "Engine").Str("method", "initialSync").Logger()
	errGroup := new(errgroup.Group)

	engineCtx := EngineCtx{
		Engine:        e,
		Ctx:           ctx,
		BroadcastData: e.indexersGate.BroadcastDataEvent,
		Limiter:       e.providerLimiter,
	}

	for _, p := range e.dataProducers {
		errGroup.Go(func() error {
			if err := p.OnProduce(engineCtx, models.DataProduceTrigger{
				StartBlock: e.cfg.NetworkConfig.StartBlock,
				EndBlock:   e.cfg.NetworkConfig.EndBlock,
			}); err != nil {
				l.Error().Err(err).Msg("data produced returned error")
				return err
			}

			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		l.Error().Err(err).Msg("some of data producers failed")
		return err
	}

	return nil
}

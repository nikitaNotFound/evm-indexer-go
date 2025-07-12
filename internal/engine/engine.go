package engine

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"github.com/nikitaNotFound/evm-indexer-go/internal/config"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
)

type Engine struct {
	dataProducers []DataProducer
	cfg           *config.Config
	indexersGate  *IndexersGate
}

type EngineCtx struct {
	Engine        *Engine
	Ctx           context.Context
	BroadcastData func(topic string, data interface{}) error
}

func CreateEngine(cfg *config.Config, dp []DataProducer) Engine {
	return Engine{
		dataProducers: dp,
		cfg:           cfg,
		indexersGate:  NewIndexersGate(),
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

func (e *Engine) initialSync(ctx context.Context) error {
	l := log.With().Str("component", "Engine").Str("method", "initialSync").Logger()
	errGroup := new(errgroup.Group)

	engineCtx := EngineCtx{
		Engine: e,
		Ctx:    ctx,
		BroadcastData: func(topic string, data interface{}) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			event := models.ProducedDataEvent{
				Data:      data,
				Timestamp: time.Now(),
				Trigger:   models.InitialTrigger,
			}

			return e.indexersGate.BroadcastDataEvent(topic, event)
		},
	}

	for _, p := range e.dataProducers {
		errGroup.Go(func() error {
			if err := p.OnProduce(engineCtx, models.DataProduceTrigger{
				TriggerType: models.InitialTrigger,
				StartBlock:  e.cfg.NetworkConfig.StartBlock,
				EndBlock:    e.cfg.NetworkConfig.EndBlock,
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

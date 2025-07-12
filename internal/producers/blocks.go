package producers

import (
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/rs/zerolog/log"
)

const TopicName = "blocks"

type BlocksProducer struct {
}

func NewBlocksProducer() *BlocksProducer {
	return &BlocksProducer{}
}

type Block struct {
	Number uint64
	Hash   string
}

func (p *BlocksProducer) OnProduce(
	e engine.EngineCtx,
	trigger models.DataProduceTrigger,
) error {
	l := log.With().Str("component", "BlocksProducer").Str("method", "OnProduce").Logger()

	l.Info().Interface("trigger", trigger).Msg("producing blocks")

	if err := e.BroadcastData(TopicName, Block{Number: 123, Hash: "0x123"}); err != nil {
		l.Error().Err(err).Msg("failed to broadcast data")
		return err
	}

	return nil
}

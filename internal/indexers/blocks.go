package indexers

import (
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/rs/zerolog/log"
)

type BlocksIndexer struct {
}

func NewBlocksIndexer() *BlocksIndexer {
	return &BlocksIndexer{}
}

func (i *BlocksIndexer) OnDataEvent(topic string, data models.ProducedDataEvent) error {
	l := log.With().Str("component", "BlocksIndexer").Str("method", "OnDataEvent").Logger()

	l.Info().Str("topic", topic).Interface("data", data).Msg("blocks indexer received data")

	return nil
}

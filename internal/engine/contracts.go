package engine

import (
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
)

type DataProducer interface {
	OnProduce(ctx EngineCtx, trigger models.DataProduceTrigger) error
}

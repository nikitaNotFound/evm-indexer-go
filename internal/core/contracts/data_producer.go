package contracts

import "github.com/nikitaNotFound/evm-indexer-go/internal/core/models"

type DataProducer interface {
	OnProduceTrigger(t models.TriggerDataProduce) error
}

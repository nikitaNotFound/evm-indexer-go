package models

type DataProduceTriggerType string

const (
	Initial   DataProduceTriggerType = "intial"
	EachBlock DataProduceTriggerType = "each_block"
)

type TriggerDataProduce struct {
	StartBlock  uint64
	EndBlock    uint64
	TriggerType DataProduceTriggerType
}

type DataProducedEvent struct {
}

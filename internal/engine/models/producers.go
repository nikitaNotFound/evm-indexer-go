package models

import (
	"time"
)

type DataProduceTriggerType string

const (
	InitialTrigger   DataProduceTriggerType = "initial"
	EachBlockTrigger DataProduceTriggerType = "each_block"
)

type DataProduceTrigger struct {
	TriggerType DataProduceTriggerType
	StartBlock  int64
	EndBlock    int64
}

type ProducedDataEvent struct {
	Data      interface{}
	Timestamp time.Time
	Trigger   DataProduceTriggerType
}

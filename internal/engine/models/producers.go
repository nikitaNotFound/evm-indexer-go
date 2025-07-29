package models

import (
	"time"
)

type DataProduceTriggerType string

type DataProduceTrigger struct {
	TriggerType DataProduceTriggerType
	StartBlock  int64
	EndBlock    int64
}

type ProducedDataEvent struct {
	Data         interface{}
	timestampUtc time.Time
}

func (e *ProducedDataEvent) TimestampUTC() time.Time {
	return e.timestampUtc
}

func NewProducedDataEvent(data interface{}) ProducedDataEvent {
	return ProducedDataEvent{
		Data:         data,
		timestampUtc: time.Now().UTC(),
	}
}

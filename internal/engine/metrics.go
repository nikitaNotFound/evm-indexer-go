package engine

import "sync/atomic"

type MetricsService struct {
	metrics map[string]atomic.Int64
}

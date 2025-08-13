package smartlim

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type MockedReceiver struct {
	mu sync.Mutex

	received int
	limit    int
}

func NewMockedReceiver(limit int) *MockedReceiver {
	return &MockedReceiver{
		limit: limit,
	}
}

func (r *MockedReceiver) Start(ctx context.Context) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			r.mu.Lock()
			fmt.Printf("Received: %d\n", r.received)
			r.received = 0
			r.mu.Unlock()
		case <-ctx.Done():
			return nil
		}
	}
}
func (r *MockedReceiver) Receive(ctx context.Context, data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.received >= r.limit {
		return fmt.Errorf("limit reached")
	}

	r.received++

	return nil
}

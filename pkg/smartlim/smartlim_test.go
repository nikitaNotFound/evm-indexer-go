package smartlim

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestLeakyBucket_Allow(t *testing.T) {
	tests := []struct {
		name          string
		ratePerSec    int
		calls         int
		expectAllowed int
	}{
		{
			name:          "Basic rate limiting",
			ratePerSec:    10,
			calls:         15,
			expectAllowed: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := StartSmartLimiter(tt.ratePerSec, tt.ratePerSec)

			allowed := 0
			for i := 0; i < tt.calls; i++ {
				if b.allow() {
					allowed++
				}
			}

			if allowed != tt.expectAllowed {
				t.Errorf("Expected %d allowed calls, got %d", tt.expectAllowed, allowed)
			}
		})
	}
}

func TestLeakyBucket_Leaking(t *testing.T) {
	b := StartSmartLimiter(10, 10) // 10 per second, capacity 10

	// Fill the bucket
	for i := 0; i < 10; i++ {
		if !b.allow() {
			t.Fatalf("Expected allow to return true for the first 10 calls")
		}
	}

	// Bucket should be full now
	if b.allow() {
		t.Fatalf("Expected allow to return false when bucket is full")
	}

	// Wait for 200ms, which should allow 2 more requests at 10/sec rate
	time.Sleep(200 * time.Millisecond)

	// We should be able to make 2 requests now
	allowed := 0
	for i := 0; i < 3; i++ {
		if b.allow() {
			allowed++
		}
	}

	// Allow for some timing uncertainty, expecting around 2
	if allowed < 1 || allowed > 3 {
		t.Errorf("Expected around 2 allowed calls after waiting, got %d", allowed)
	}
}

func TestProcess_DirectExecution(t *testing.T) {
	b := StartSmartLimiter(10, 10)
	ctx := context.Background()

	result, err := Process(b, ctx, func() (string, error) {
		return "success", nil
	})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != "success" {
		t.Errorf("Expected result to be 'success', got '%s'", result)
	}

	// Fill the bucket
	for i := 0; i < 10; i++ {
		b.allow()
	}

	// Set a small timeout to avoid long test
	b.SetQueueTimeout(50 * time.Millisecond)

	// This should time out since the bucket is full and queue processing isn't started
	_, err = Process(b, ctx, func() (string, error) {
		return "should not execute", nil
	})

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Expected deadline exceeded error, got %v", err)
	}
}

func TestProcess_WithQueue(t *testing.T) {
	b := StartSmartLimiter(100, 1) // 100 per second, capacity 1
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Use up the one available slot
	if !b.allow() {
		t.Fatal("First call should be allowed")
	}

	// Create a channel to signal when our test function has executed
	executed := make(chan struct{})

	// This will be queued
	go func() {
		result, err := Process(b, ctx, func() (string, error) {
			close(executed)
			return "queued success", nil
		})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if result != "queued success" {
			t.Errorf("Expected result to be 'queued success', got '%s'", result)
		}
	}()

	// Wait for the queued function to execute or time out
	select {
	case <-executed:
		// Success - function executed via queue
	case <-time.After(200 * time.Millisecond):
		t.Fatal("Timed out waiting for queued function to execute")
	}
}

func TestQueueLength(t *testing.T) {
	b := StartSmartLimiter(10, 10)

	if b.QueueLength() != 0 {
		t.Errorf("Expected empty queue, got length %d", b.QueueLength())
	}

	// Enqueue some items
	for i := 0; i < 5; i++ {
		b.enqueueAllowance()
	}

	if b.QueueLength() != 5 {
		t.Errorf("Expected queue length 5, got %d", b.QueueLength())
	}
}

func TestConcurrentUsage(t *testing.T) {
	b := StartSmartLimiter(50, 50)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	successCount := 0
	var mu sync.Mutex

	// Launch 100 concurrent requests
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			result, err := Process(b, ctx, func() (bool, error) {
				return true, nil
			})

			if err == nil && result {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// We should have at least 10 successes (initial capacity)
	// and more as the bucket leaks
	if successCount < 10 {
		t.Errorf("Expected at least 10 successful calls, got %d", successCount)
	}

	t.Logf("Concurrent test had %d successful calls out of 100", successCount)
}

func TestContextCancellation(t *testing.T) {
	b := StartSmartLimiter(1, 1) // Very slow rate to ensure queueing

	// Use up the capacity
	if !b.allow() {
		t.Fatal("First call should be allowed")
	}

	// Create a context we can cancel
	ctx, cancel := context.WithCancel(context.Background())

	// Create a channel to know when the function completes
	done := make(chan struct{})

	// Start a process that will be queued
	go func() {
		defer close(done)
		_, err := Process(b, ctx, func() (bool, error) {
			return true, nil
		})

		// Should get cancelled before executing
		if !errors.Is(err, context.Canceled) {
			t.Errorf("Expected context.Canceled error, got %v", err)
		}
	}()

	// Wait a moment to ensure it's queued
	time.Sleep(50 * time.Millisecond)

	// Cancel the context
	cancel()

	// Wait for the function to complete
	select {
	case <-done:
		// Success - function returned with context cancellation
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Timed out waiting for function to handle cancellation")
	}
}

// Benchmarks

// BenchmarkAllow measures the performance of the allow() method
func BenchmarkAllow(b *testing.B) {
	bucket := StartSmartLimiter(1000, 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bucket.allow()
	}
}

// BenchmarkProcess measures the performance of the Process function for direct execution
func BenchmarkProcess(b *testing.B) {
	bucket := StartSmartLimiter(1000, 1000) // Ensure capacity for all iterations
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Process(bucket, ctx, func() (int, error) {
			return 1, nil
		})
	}
}

// BenchmarkProcessWithContentionLow measures Process performance with 10% contention
func BenchmarkProcessWithContentionLow(b *testing.B) {
	bucket := StartSmartLimiter(1000, 1000) // 10% of iterations can succeed immediately
	ctx := context.Background()
	bucket.SetQueueTimeout(1 * time.Millisecond) // Short timeout for failing quickly

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Process(bucket, ctx, func() (int, error) {
			return 1, nil
		})
	}
}

// BenchmarkProcessWithContentionHigh measures Process performance with 90% contention
func BenchmarkProcessWithContentionHigh(b *testing.B) {
	bucket := StartSmartLimiter(1000, 1000) // 10% of iterations can succeed immediately
	ctx := context.Background()
	bucket.SetQueueTimeout(1 * time.Millisecond) // Short timeout for failing quickly

	// Use up 90% of the capacity
	for i := 0; i < (b.N/10)*9; i++ {
		bucket.allow()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Process(bucket, ctx, func() (int, error) {
			return 1, nil
		})
	}
}

// BenchmarkConcurrentProcess measures the performance under concurrent load
func BenchmarkConcurrentProcess(b *testing.B) {
	bucket := StartSmartLimiter(1000, 1000)
	ctx := context.Background()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Process(bucket, ctx, func() (int, error) {
				return 1, nil
			})
		}
	})
}

package smartlim

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type SmartLimiter struct {
	maxBurst           int
	lastLeakTime       time.Time
	leakInterval       time.Duration
	maxRps             int
	lastTokensFillTime time.Time
	tokens             int
	water              int
	mu                 sync.Mutex

	allowQueue *Queue[chan struct{}]

	queueStart sync.Once

	queueTimeout time.Duration

	channelPool *sync.Pool
}

const (
	DefaultQueueSize = 512
)

var (
	ErrRateLimitExceeded = errors.New("rate limit exceeded")
	ErrQueueTimeout      = errors.New("queue timeout")
)

func StartSmartLimiter(ratePerSec int, maxBurst int) *SmartLimiter {
	limiter := &SmartLimiter{
		maxBurst:           maxBurst,
		maxRps:             ratePerSec,
		tokens:             maxBurst,
		lastTokensFillTime: time.Now(),
		lastLeakTime:       time.Now(),
		leakInterval:       time.Second / time.Duration(ratePerSec),
		allowQueue:         NewQueue[chan struct{}](),
		queueStart:         sync.Once{},
		queueTimeout:       time.Duration(0),
		channelPool: &sync.Pool{
			New: func() interface{} {
				return make(chan struct{}, 1)
			},
		},
	}

	limiter.startQueueProcessing(context.Background())

	return limiter
}

func (b *SmartLimiter) SetQueueTimeout(timeout time.Duration) {
	b.queueTimeout = timeout
}

func (b *SmartLimiter) startQueueProcessing(ctx context.Context) {
	b.queueStart.Do(func() {
		go func() {
			ticker := time.NewTicker(b.leakInterval)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					if el, release := b.allowQueue.BlockingProcess(); el != nil {
						if b.allow() {
							el <- struct{}{}
						} else {
							for {
								if b.allow() {
									el <- struct{}{}
									break
								}
							}
						}
						release()
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	})
}

func Process[T any](b *SmartLimiter, ctx context.Context, next func() (T, error)) (T, error) {
	if b.QueueLength() != 0 {
		return processWait(b, ctx, next)
	}

	if b.allow() {
		return next()
	}

	return processWait(b, ctx, next)
}

func processWait[T any](b *SmartLimiter, ctx context.Context, next func() (T, error)) (T, error) {
	notifyChan, id := b.enqueueAllowance()
	defer b.channelPool.Put(notifyChan)

	if b.queueTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, b.queueTimeout)
		defer cancel()
	}

	var zero T

	select {
	case <-notifyChan:
		return next()
	case <-ctx.Done():
		b.allowQueue.DequeueByID(id)
		_, _ = <-notifyChan // Drain the channel
		b.channelPool.Put(notifyChan)
		return zero, ctx.Err()
	}
}

func (b *SmartLimiter) QueueLength() int {
	return b.allowQueue.Len()
}

func (b *SmartLimiter) enqueueAllowance() (chan struct{}, uuid.UUID) {
	notifyChan := b.channelPool.Get().(chan struct{})
	if notifyChan == nil {
		notifyChan = make(chan struct{}, 1)
	}

	id := b.allowQueue.Enqueue(notifyChan)
	return notifyChan, id
}

// cancelAllowance cancels a specific allowance request and returns the channel to the pool
// Returns true if the request was successfully canceled
func (b *SmartLimiter) cancelAllowance(ch chan struct{}) bool {
	// Check if the channel already has a value (was already processed)
	select {
	case <-ch:
		// Channel was already processed, can't cancel
		return false
	default:
		// Try to add cancellation signal
		select {
		case ch <- struct{}{}:
			// Successfully marked as canceled
			// Empty the channel immediately
			<-ch

			// Return the channel to the pool for reuse
			b.channelPool.Put(ch)
			return true
		default:
			// Channel was processed by another goroutine just now
			return false
		}
	}
}

func (b *SmartLimiter) allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	fillElapsed := now.Sub(b.lastTokensFillTime)
	if fillElapsed > time.Second {
		b.lastTokensFillTime = now
		b.tokens = b.maxRps
		fillElapsed = time.Duration(0)
	}

	if b.tokens <= 0 {
		return false
	}

	// Can be used to maximize number of usages, but can lead to big burst
	// leakInterval := (time.Second - fillElapsed) / time.Duration(b.tokens)

	elapsed := now.Sub(b.lastLeakTime)
	leaked := int(elapsed / b.leakInterval)

	if leaked > 0 {
		b.water = max(0, b.water-leaked)
		b.lastLeakTime = now
	}

	if b.water < b.maxBurst {
		b.water++
		b.tokens--
		return true
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

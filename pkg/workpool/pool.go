package workpool

import (
	"context"
	"sync"
)

type WorkPool[T any] struct {
	workers int
	options *options

	poolCtx context.Context
	cancel  context.CancelFunc

	workChan chan func() (T, error)
	wg       sync.WaitGroup
	errCh    chan error
	resCh    chan T
}

type options struct {
	ctx context.Context
}

func defaultOptions() *options {
	return &options{
		ctx: context.Background(),
	}
}

type Option func(*options)

func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

func NewWorkPool[T any](workersCount int, opts ...Option) *WorkPool[T] {
	opt := defaultOptions()
	for _, o := range opts {
		o(opt)
	}

	poolCtx, cancel := context.WithCancel(opt.ctx)

	p := &WorkPool[T]{
		workers:  workersCount,
		options:  opt,
		poolCtx:  poolCtx,
		cancel:   cancel,
		workChan: make(chan func() (T, error)),
		errCh:    make(chan error),
		resCh:    make(chan T),
	}

	for i := 0; i < workersCount; i++ {
		go p.startWorker(opt.ctx)
	}

	return p
}

// Do is used to submit work to the pool
func (p *WorkPool[T]) Do(work func() (T, error)) {
	p.wg.Add(1)
	p.workChan <- work
}

// Wait is used to wait for all work to be completed
func (p *WorkPool[T]) Wait() {
	p.wg.Wait()
}

// Stop is used to stop the pool
func (p *WorkPool[T]) Stop() {
	p.cancel()
	close(p.workChan)
	close(p.errCh)
	close(p.resCh)
}

// GetResults is used to get the results of the work
func (p *WorkPool[T]) Results() <-chan T {
	return p.resCh
}

// Errors is used to get the errors of the work
func (p *WorkPool[T]) Errors() <-chan error {
	return p.errCh
}

func (p *WorkPool[T]) startWorker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case work := <-p.workChan:
			p.processWork(work)
		}
	}
}

func (p *WorkPool[T]) processWork(work func() (T, error)) {
	defer p.wg.Done()

	res, err := work()
	if err != nil {
		p.errCh <- err
	} else {
		p.resCh <- res
	}
}

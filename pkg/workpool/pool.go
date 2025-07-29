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
	wg       *sync.WaitGroup
	errCh    chan error
	resCh    chan T
}

type options struct {
	ctx         context.Context
	errChanSize int
	resChanSize int
	poolSize    int
}

func defaultOptions() *options {
	return &options{
		ctx:         context.Background(),
		errChanSize: 512,
		resChanSize: 512,
		poolSize:    512,
	}
}

type Option func(*options)

func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

func WithErrChanSize(size int) Option {
	return func(o *options) {
		o.errChanSize = size
	}
}

func WithResChanSize(size int) Option {
	return func(o *options) {
		o.resChanSize = size
	}
}

func WithPoolSize(size int) Option {
	return func(o *options) {
		o.poolSize = size
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
		workChan: make(chan func() (T, error), opt.poolSize),
		errCh:    make(chan error, opt.errChanSize),
		resCh:    make(chan T, opt.resChanSize),
		wg:       &sync.WaitGroup{},
	}

	for range workersCount {
		go p.startWorker()
	}

	return p
}

// Enqueue is used to submit work to the pool, potential blocking if the pool is full
func (p *WorkPool[T]) Enqueue(work func() (T, error)) {
	p.wg.Add(1)
	p.workChan <- work
}

// WaitAndStop is used to wait for all work to be completed and then stop the pool
func (p *WorkPool[T]) WaitAndStop() {
	p.wg.Wait()
	p.cancel()
	close(p.workChan)
	close(p.errCh)
	close(p.resCh)
}

// InstantStop is used to stop the pool immediately
func (p *WorkPool[T]) InstantStop() {
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

func (p *WorkPool[T]) startWorker() {
	for {
		select {
		case <-p.poolCtx.Done():
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

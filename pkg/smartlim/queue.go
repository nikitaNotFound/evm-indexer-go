package smartlim

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type QueueEl[T any] struct {
	ID    uuid.UUID
	Value T
	Next  *QueueEl[T]
	Prev  *QueueEl[T]
}

type Queue[T any] struct {
	Head *QueueEl[T]
	Tail *QueueEl[T]

	hashMap map[uuid.UUID]*QueueEl[T]

	Size int

	mu sync.Mutex
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		hashMap: make(map[uuid.UUID]*QueueEl[T]),
		Size:    0,
		Head:    nil,
		Tail:    nil,
		mu:      sync.Mutex{},
	}
}

func (q *Queue[T]) Len() int {
	return q.Size
}

func (q *Queue[T]) BlockingProcess() (T, func()) {
	var zero T
	if q.Head == nil {
		return zero, func() {}
	}

	q.mu.Lock()

	released := func() {
		defer q.mu.Unlock()

		q.dequeueByID(q.Head.ID)
	}

	return q.Head.Value, released
}

func (q *Queue[T]) Enqueue(value T) uuid.UUID {
	q.mu.Lock()
	defer q.mu.Unlock()

	id := uuid.New()
	el := &QueueEl[T]{
		ID:    id,
		Value: value,
		Next:  nil,
		Prev:  nil,
	}

	if q.Size == 0 {
		q.Head = el
		q.Tail = el
	} else {
		q.Tail.Next = el
		el.Prev = q.Tail
		q.Tail = el
	}

	q.hashMap[id] = el
	q.Size++

	return id
}

func (q *Queue[T]) DequeueByID(id uuid.UUID) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.dequeueByID(id)
}

// Should not be used without mutex lock
func (q *Queue[T]) dequeueByID(id uuid.UUID) error {
	el, ok := q.hashMap[id]
	if !ok {
		return fmt.Errorf("element not found")
	}

	if el.Prev != nil {
		el.Prev.Next = el.Next
	} else {
		q.Head = el.Next
	}

	if el.Next != nil {
		el.Next.Prev = el.Prev
	} else {
		q.Tail = el.Prev
	}

	delete(q.hashMap, id)
	q.Size--

	return nil
}

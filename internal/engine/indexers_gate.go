package engine

import (
	"context"
	"errors"
	"sync"

	"github.com/nikitaNotFound/evm-indexer-go/internal/engine/models"
	"github.com/rs/zerolog/log"
)

var (
	ErrTopicAlreadyExists = errors.New("topic already exists")
	ErrTopicNotFound      = errors.New("topic not found")
)

type TopicSubscriber interface {
	OnDataEvent(ctx context.Context, topic string, event models.ProducedDataEvent) error
}

type IndexersGate struct {
	topicSubs map[string][]TopicSubscriber
	ctx       context.Context
	cancel    context.CancelFunc

	mu              sync.RWMutex
	consumeEventsWG sync.WaitGroup
}

// NewIndexersGate creates a new indexers gate
func NewIndexersGate() *IndexersGate {
	ctx, cancel := context.WithCancel(context.Background())
	return &IndexersGate{
		topicSubs:       map[string][]TopicSubscriber{},
		consumeEventsWG: sync.WaitGroup{},
		mu:              sync.RWMutex{},
		ctx:             ctx,
		cancel:          cancel,
	}
}

// Stop stops the indexers gate
func (g *IndexersGate) Stop() {
	g.cancel()
	g.consumeEventsWG.Wait()
}

// WaitFinish waits for all events to be consumed
func (g *IndexersGate) WaitFinish() {
	g.consumeEventsWG.Wait()
}

func (g *IndexersGate) topicExists(topic string) bool {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.topicSubs[topic] != nil
}

// CreateTopic creates a new topic
func (g *IndexersGate) CreateTopic(topic string) error {
	if g.topicExists(topic) {
		return ErrTopicAlreadyExists
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	g.topicSubs[topic] = make([]TopicSubscriber, 0)

	return nil
}

// Subscribe subscribes a new subscriber to a topic
func (g *IndexersGate) Subscribe(topic string, subscriber TopicSubscriber) error {
	if !g.topicExists(topic) {
		return ErrTopicNotFound
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	g.topicSubs[topic] = append(g.topicSubs[topic], subscriber)

	return nil
}

// BroadcastDataEvent broadcasts a data event to all subscribers of a topic
func (g *IndexersGate) BroadcastDataEvent(topic string, data models.ProducedDataEvent) error {
	if !g.topicExists(topic) {
		return ErrTopicNotFound
	}

	for _, subscriber := range g.topicSubs[topic] {
		g.consumeEventsWG.Add(1)
		go func(subscriber TopicSubscriber) {
			defer g.consumeEventsWG.Done()

			if err := subscriber.OnDataEvent(g.ctx, topic, data); err != nil {
				log.Error().Err(err).Msg("failed to broadcast data event")
			}
		}(subscriber)
	}

	return nil
}

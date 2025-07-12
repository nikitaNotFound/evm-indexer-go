package engine

import (
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
	OnDataEvent(topic string, event models.ProducedDataEvent) error
}

type IndexersGate struct {
	topicSubs map[string][]TopicSubscriber

	mu sync.RWMutex

	wg sync.WaitGroup
}

func NewIndexersGate() *IndexersGate {
	return &IndexersGate{
		topicSubs: map[string][]TopicSubscriber{},
		wg:        sync.WaitGroup{},
		mu:        sync.RWMutex{},
	}
}

func (g *IndexersGate) WaitFinish() {
	g.wg.Wait()
}

func (g *IndexersGate) topicExists(topic string) bool {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.topicSubs[topic] != nil
}

func (g *IndexersGate) CreateTopic(topic string) error {
	if g.topicExists(topic) {
		return ErrTopicAlreadyExists
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	g.topicSubs[topic] = make([]TopicSubscriber, 0)

	return nil
}

func (g *IndexersGate) Subscribe(topic string, subscriber TopicSubscriber) error {
	if !g.topicExists(topic) {
		return ErrTopicNotFound
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	g.topicSubs[topic] = append(g.topicSubs[topic], subscriber)

	return nil
}

func (g *IndexersGate) BroadcastDataEvent(topic string, data models.ProducedDataEvent) error {
	if !g.topicExists(topic) {
		return ErrTopicNotFound
	}

	for _, subscriber := range g.topicSubs[topic] {
		g.wg.Add(1)
		go func(subscriber TopicSubscriber) {
			defer g.wg.Done()

			if err := subscriber.OnDataEvent(topic, data); err != nil {
				log.Error().Err(err).Msg("failed to broadcast data event")
			}
		}(subscriber)
	}

	return nil
}

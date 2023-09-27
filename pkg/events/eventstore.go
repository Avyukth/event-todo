package events

import (
	"errors"
	"sync"
)

var (
	ErrEventNotFound = errors.New("event not found")
)



type EventStore struct {
	mu      sync.RWMutex
	storage map[string][]Event
}

func NewInMemoryEventStore() *EventStore {
	return &EventStore{
		mu : sync.RWMutex{},
		storage: make(map[string][]Event),
	}
}

func (es *EventStore) Save(aggregateID string, event Event) error {
	es.mu.Lock()
	defer es.mu.Unlock()

	es.storage[aggregateID] = append(es.storage[aggregateID], event)

	return nil
}

func (es *EventStore) Load(aggregateID string) ([]Event, error) {
	es.mu.RLock()
	defer es.mu.RUnlock()

	events, ok := es.storage[aggregateID]
	if !ok {
		return nil, ErrEventNotFound
	}

	return events, nil
}

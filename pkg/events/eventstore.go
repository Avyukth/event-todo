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
	storage map[string][]Event // map[aggregateID]events
}

func NewInMemoryEventStore() *EventStore {
	return &EventStore{
		storage: make(map[string][]Event),
	}
}

func (es *EventStore) Save(aggregateID string, event Event) error {
	es.mu.Lock()
	defer es.mu.Unlock()

	// Append the new event to the list of events for the given aggregateID
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

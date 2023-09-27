package repo

import (
	ev "event-todo/pkg/events"
	"sync"
)




type EventStore struct {
	mu      sync.RWMutex
	storage map[string][]ev.Event
}

func NewInMemoryEventStore() *EventStore {
	return &EventStore{
		mu : sync.RWMutex{},
		storage: make(map[string][]ev.Event),
	}
}

func (es *EventStore) Save(aggregateID string, event ev.Event) error {
	es.mu.Lock()
	defer es.mu.Unlock()

	es.storage[aggregateID] = append(es.storage[aggregateID], event)

	return nil
}

func (es *EventStore) Load(aggregateID string) ([]ev.Event, error) {
	es.mu.RLock()
	defer es.mu.RUnlock()

	events, ok := es.storage[aggregateID]
	if !ok {
		return nil, ev.ErrEventNotFound
	}

	return events, nil
}

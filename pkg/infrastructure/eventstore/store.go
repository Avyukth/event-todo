package eventstore

import (
	"errors"
	"sync"

	"event-todo/pkg/domains"
)

// EventStore is an interface defining the methods that an event store should have.
type EventStore interface {
	SaveEvents(aggregateID string, events []domain.Event) error
	GetEventsForAggregate(aggregateID string) ([]domain.Event, error)
}

// InMemoryEventStore is a simple, in-memory implementation of the EventStore interface.
type InMemoryEventStore struct {
	mu      sync.RWMutex
	storage map[string][]domain.Event // map[aggregateID]events
}

// NewInMemoryEventStore initializes a new InMemoryEventStore.
func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		storage: make(map[string][]domain.Event),
	}
}

// SaveEvents stores the provided events for the given aggregateID.
func (s *InMemoryEventStore) SaveEvents(aggregateID string, events []domain.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.storage[aggregateID]; !ok {
		s.storage[aggregateID] = make([]domain.Event, 0)
	}

	s.storage[aggregateID] = append(s.storage[aggregateID], events...)
	return nil
}

// GetEventsForAggregate retrieves all events for the given aggregateID.
func (s *InMemoryEventStore) GetEventsForAggregate(aggregateID string) ([]domain.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if events, ok := s.storage[aggregateID]; ok {
		return events, nil
	}

	return nil, errors.New("no events found for the given aggregateID")
}

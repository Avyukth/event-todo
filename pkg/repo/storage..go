package repo

import (
	ev "event-todo/pkg/events"
)


type EventStorer interface {
	Save(aggregateID string, event ev.Event) error
	Load(aggregateID string) ([]ev.Event, error)
}

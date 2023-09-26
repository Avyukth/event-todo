package events

import (
	"errors"
)

// Define errors
var (
	ErrInvalidEventType = errors.New("invalid event type")
)

// Event is the interface that all events must implement.
type Event interface {
	EventType() string
}

// Aggregate is the interface that all aggregate roots must implement.
type Aggregate interface {
	ApplyEvent(event Event) error
}

// TaskCreatedEvent is emitted when a new task is created.
type TaskCreatedEvent struct {
	ID    string
	Title string
}

func (e *TaskCreatedEvent) EventType() string {
	return "TaskCreated"
}

// TaskCompletedEvent is emitted when a task is completed.
type TaskCompletedEvent struct {
	ID string
}

func (e *TaskCompletedEvent) EventType() string {
	return "TaskCompleted"
}

// TaskDeletedEvent is emitted when a task is deleted.
type TaskDeletedEvent struct {
	ID string
}

func (e *TaskDeletedEvent) EventType() string {
	return "TaskDeleted"
}

// ApplyEventToAggregate applies an event to an aggregate.
func ApplyEventToAggregate(aggregate Aggregate, event Event) error {
	switch e := event.(type) {
	case *TaskCreatedEvent:
		return aggregate.(*todo.TaskAggregate).ApplyTaskCreatedEvent(e)
	case *TaskCompletedEvent:
		return aggregate.(*todo.TaskAggregate).ApplyTaskCompletedEvent(e)
	case *TaskDeletedEvent:
		return aggregate.(*todo.TaskAggregate).ApplyTaskDeletedEvent(e)
	default:
		return ErrInvalidEventType
	}
}

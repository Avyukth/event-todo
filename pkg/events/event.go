package events

import (
	"errors"
)


var (
	ErrInvalidEventType = errors.New("invalid event type")

	ErrEventNotFound = errors.New("event not found")
)


type Event interface {
	EventType() string
}


type Aggregate interface {
	ApplyEvent(event Event) error
}


type TaskCreatedEvent struct {
	ID    string
	Title string
}

func (e *TaskCreatedEvent) EventType() string {
	return "TaskCreated"
}


type TaskCompletedEvent struct {
	ID string
}

func (e *TaskCompletedEvent) EventType() string {
	return "TaskCompleted"
}


type TaskDeletedEvent struct {
	ID string
}

func (e *TaskDeletedEvent) EventType() string {
	return "TaskDeleted"
}


func ApplyEventToAggregate(aggregate Aggregate, event Event) error {
	switch e := event.(type) {
	case *TaskCreatedEvent:
		return aggregate.(*TaskAggregate).ApplyTaskCreatedEvent(e)
	case *TaskCompletedEvent:
		return aggregate.(*TaskAggregate).ApplyTaskCompletedEvent(e)
	case *TaskDeletedEvent:
		return aggregate.(*TaskAggregate).ApplyTaskDeletedEvent(e)
	default:
		return ErrInvalidEventType
	}
}

package todo

import (
	"errors"
	"event-todo/pkg/events"
)

// Define errors
var (
	ErrTaskAlreadyCompleted = errors.New("task already completed")
	ErrTaskAlreadyDeleted   = errors.New("task already deleted")
)

// TaskAggregate represents the aggregate root for a task.
type TaskAggregate struct {
	ID        string
	Title     string
	Completed bool
	Deleted   bool
}

// NewTaskAggregate initializes a new TaskAggregate.
func NewTaskAggregate(id string) *TaskAggregate {
	return &TaskAggregate{
		ID: id,
	}
}

// ApplyEvent applies an event to the TaskAggregate.
func (t *TaskAggregate) ApplyEvent(event events.Event) error {
	switch e := event.(type) {
	case *events.TaskCreatedEvent:
		return t.ApplyTaskCreatedEvent(e)
	case *events.TaskCompletedEvent:
		return t.ApplyTaskCompletedEvent(e)
	case *events.TaskDeletedEvent:
		return t.ApplyTaskDeletedEvent(e)
	default:
		return events.ErrInvalidEventType
	}
}

// ApplyTaskCreatedEvent applies a TaskCreatedEvent to the TaskAggregate.
func (t *TaskAggregate) ApplyTaskCreatedEvent(event *events.TaskCreatedEvent) error {
	t.ID = event.ID
	t.Title = event.Title
	return nil
}

// ApplyTaskCompletedEvent applies a TaskCompletedEvent to the TaskAggregate.
func (t *TaskAggregate) ApplyTaskCompletedEvent(event *events.TaskCompletedEvent) error {
	if t.Completed {
		return ErrTaskAlreadyCompleted
	}
	if t.Deleted {
		return ErrTaskAlreadyDeleted
	}
	t.Completed = true
	return nil
}

// ApplyTaskDeletedEvent applies a TaskDeletedEvent to the TaskAggregate.
func (t *TaskAggregate) ApplyTaskDeletedEvent(event *events.TaskDeletedEvent) error {
	if t.Deleted {
		return ErrTaskAlreadyDeleted
	}
	t.Deleted = true
	return nil
}

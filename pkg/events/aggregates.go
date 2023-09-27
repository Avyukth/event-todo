package events

import (
	"errors"
)


var (
	ErrTaskAlreadyCompleted = errors.New("task already completed")
	ErrTaskAlreadyDeleted   = errors.New("task already deleted")
)


type TaskAggregate struct {
	ID        string
	Title     string
	Completed bool
	Deleted   bool
}


func NewTaskAggregate(id string) *TaskAggregate {
	return &TaskAggregate{
		ID: id,
	}
}


func (t *TaskAggregate) ApplyEvent(event Event) error {
	switch e := event.(type) {
	case *TaskCreatedEvent:
		return t.ApplyTaskCreatedEvent(e)
	case *TaskCompletedEvent:
		return t.ApplyTaskCompletedEvent(e)
	case *TaskDeletedEvent:
		return t.ApplyTaskDeletedEvent(e)
	default:
		return ErrInvalidEventType
	}
}


func (t *TaskAggregate) ApplyTaskCreatedEvent(event *TaskCreatedEvent) error {
	t.ID = event.ID
	t.Title = event.Title
	return nil
}


func (t *TaskAggregate) ApplyTaskCompletedEvent(event *TaskCompletedEvent) error {
	if t.Completed {
		return ErrTaskAlreadyCompleted
	}
	if t.Deleted {
		return ErrTaskAlreadyDeleted
	}
	t.Completed = true
	return nil
}


func (t *TaskAggregate) ApplyTaskDeletedEvent(event *TaskDeletedEvent) error {
	if t.Deleted {
		return ErrTaskAlreadyDeleted
	}
	t.Deleted = true
	return nil
}

package todo

import (
	"errors"
	"event-todo/pkg/events"
	"event-todo/pkg/repo"
)


var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrTitleRequired = errors.New("title is required")
)


type CommandHandler struct {
	EventStore        *repo.EventStore
	ProjectionManager *events.ProjectionManager
}


type Command interface {
	Execute(handler *CommandHandler, aggregateID string) (events.Event, error)
}


type CreateTaskCommand struct {
	Title string
}


func (c *CreateTaskCommand) Execute(handler *CommandHandler, aggregateID string) (events.Event, error) {
	if c.Title == "" {
		return nil, ErrTitleRequired
	}

	return &events.TaskCreatedEvent{
		ID:    aggregateID,
		Title: c.Title,
	}, nil
}


type CompleteTaskCommand struct {
	ID string
}


func (c *CompleteTaskCommand) Execute(handler *CommandHandler, aggregateID string) (events.Event, error) {

	if _, err := handler.ProjectionManager.GetTask(aggregateID); err != nil {
		return nil, ErrTaskNotFound
	}


	return &events.TaskCompletedEvent{
		ID: aggregateID,
	}, nil
}


type DeleteTaskCommand struct {
	ID string
}


func (c *DeleteTaskCommand) Execute(handler *CommandHandler, aggregateID string) (events.Event, error) {
	if _, err := handler.ProjectionManager.GetTask(aggregateID); err != nil {
		return nil, ErrTaskNotFound
	}

	return &events.TaskDeletedEvent{
		ID: aggregateID,
	}, nil
}

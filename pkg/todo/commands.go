package todo

import (
	"errors"
	"event-todo/pkg/events"
)

// Define errors
var (
	ErrTaskNotFound = errors.New("task not found")
)

// CommandHandler handles commands and emits events.
type CommandHandler struct {
	EventStore events.EventStore
}

// Command represents a command that can be handled by a CommandHandler.
type Command interface {
	Execute(handler *CommandHandler, aggregateID string) (events.Event, error)
}

// CreateTaskCommand creates a new task.
type CreateTaskCommand struct {
	Title string
}

// Execute executes the CreateTaskCommand.
func (c *CreateTaskCommand) Execute(handler *CommandHandler, aggregateID string) (events.Event, error) {
	// Validate the command parameters
	if c.Title == "" {
		return nil, errors.New("title is required")
	}

	// Emit a TaskCreatedEvent
	return &events.TaskCreatedEvent{
		ID:    aggregateID,
		Title: c.Title,
	}, nil
}

// CompleteTaskCommand completes an existing task.
type CompleteTaskCommand struct{
    ID string
}

// Execute executes the CompleteTaskCommand.
func (c *CompleteTaskCommand) Execute(handler *CommandHandler, aggregateID string) (events.Event, error) {
	// Check if the task exists
	_, err := handler.EventStore.Load(aggregateID)
	if err != nil {
		if errors.Is(err, events.ErrEventNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}

	// Emit a TaskCompletedEvent
	return &events.TaskCompletedEvent{
		ID: aggregateID,
	}, nil
}

// DeleteTaskCommand deletes an existing task.
type DeleteTaskCommand struct{}

// Execute executes the DeleteTaskCommand.
func (c *DeleteTaskCommand) Execute(handler *CommandHandler, aggregateID string) (events.Event, error) {
	// Check if the task exists
	_, err := handler.EventStore.Load(aggregateID)
	if err != nil {
		if errors.Is(err, events.ErrEventNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}

	// Emit a TaskDeletedEvent
	return &events.TaskDeletedEvent{
		ID: aggregateID,
	}, nil
}

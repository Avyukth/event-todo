package todo

import (
	"errors"
	"event-todo/pkg/events"
)

// Define errors
var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrTitleRequired = errors.New("title is required")
)

// CommandHandler handles commands and emits events.
type CommandHandler struct {
	EventStore        *events.EventStore
	ProjectionManager *ProjectionManager // Injecting ProjectionManager
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
		return nil, ErrTitleRequired
	}

	// Emit a TaskCreatedEvent
	return &events.TaskCreatedEvent{
		ID:    aggregateID,
		Title: c.Title,
	}, nil
}

// CompleteTaskCommand completes an existing task.
type CompleteTaskCommand struct {
	ID string
}

// Execute executes the CompleteTaskCommand.
func (c *CompleteTaskCommand) Execute(handler *CommandHandler, aggregateID string) (events.Event, error) {
	// Check if the task exists using ProjectionManager
	if _, err := handler.ProjectionManager.GetTask(aggregateID); err != nil {
		return nil, ErrTaskNotFound
	}

	// Emit a TaskCompletedEvent
	return &events.TaskCompletedEvent{
		ID: aggregateID,
	}, nil
}

// DeleteTaskCommand deletes an existing task.
type DeleteTaskCommand struct {
	ID string
}

// Execute executes the DeleteTaskCommand.
func (c *DeleteTaskCommand) Execute(handler *CommandHandler, aggregateID string) (events.Event, error) {
	// Check if the task exists using ProjectionManager
	if _, err := handler.ProjectionManager.GetTask(aggregateID); err != nil {
		return nil, ErrTaskNotFound
	}

	// Emit a TaskDeletedEvent
	return &events.TaskDeletedEvent{
		ID: aggregateID,
	}, nil
}

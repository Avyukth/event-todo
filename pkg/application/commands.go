package application

import (
	"errors"

	"event-todo/pkg/domains"
)

// CommandHandlers struct holds the dependencies required by command handlers
type CommandHandlers struct {
	EventStore domain.Event
}

// NewCommandHandlers initializes the command handlers with the given event store
func NewCommandHandlers(es domain.Event) *CommandHandlers {
	return &CommandHandlers{
		EventStore: es,
	}
}

// HandleCreateTodoCommand handles a command to create a new Todo item
func (h *CommandHandlers) HandleCreateTodoCommand(command *domain.TodoCreated) error {
	// Create a new Todo aggregate
	todo, err := domain.NewTodo()
	if err != nil {
		return err
	}

	// Call the appropriate method of the aggregate to handle the command
	if err := todo.Create(command); err != nil {
		return err
	}

	// Save the resulting events to the event store
	if err := h.EventStore.SaveEvents(todo.ID(), todo.Changes()); err != nil {
		return errors.New("failed to save events")
	}

	return nil
}

// Other command handlers can be defined here, for example, UpdateTodoCommand, DeleteTodoCommand, etc.

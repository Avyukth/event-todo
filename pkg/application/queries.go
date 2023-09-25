package application

import (
	"errors"

	"event-todo/pkg/domains"
)

// QueryHandlers struct holds the dependencies required by query handlers
type QueryHandlers struct {
	ReadModel domain.ReadModel
}

// NewQueryHandlers initializes the query handlers with the given read model
func NewQueryHandlers(rm domain.ReadModel) *QueryHandlers {
	return &QueryHandlers{
		ReadModel: rm,
	}
}

// HandleGetTodoQuery handles a query to get a Todo item by ID
func (h *QueryHandlers) HandleGetTodoQuery(query *domain.GetTodoQuery) (*domain.TodoProjection, error) {
	// Query the read model to get the TodoProjection
	todo, err := h.ReadModel.GetTodoByID(query.ID)
	if err != nil {
		return nil, errors.New("failed to get todo")
	}

	return todo, nil
}

// HandleGetAllTodosQuery handles a query to get all Todo items
func (h *QueryHandlers) HandleGetAllTodosQuery(query *domain.GetAllTodosQuery) ([]*domain.TodoProjection, error) {
	// Query the read model to get all TodoProjections
	todos, err := h.ReadModel.GetAllTodos()
	if err != nil {
		return nil, errors.New("failed to get todos")
	}

	return todos, nil
}

// Other query handlers can be defined here, for example, to get Todo items by status, by user, etc.

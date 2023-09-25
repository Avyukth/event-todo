package domain

import (
	"errors"
	"time"
)

// Todo is the aggregate root and domain model representing a Todo item
type Todo struct {
	ID          string    // Unique identifier for the Todo item
	Title       string    // Title of the Todo item
	Description string    // Description of the Todo item
	Status      string    // Status of the Todo item
	CreatedAt   time.Time // Timestamp representing when the Todo item was created
	CompletedAt *time.Time // Timestamp representing when the Todo item was completed, nil if not completed
}

// NewTodo creates a new Todo item and generates the TodoCreated event
func NewTodo(id string, title string, description string) (*Todo, error) {
	if id == "" || title == "" {
		return nil, errors.New("id and title cannot be empty")
	}
	createdAt := time.Now()
	todo := &Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      "created",
		CreatedAt:   createdAt,
	}
	return todo, nil
}

// Complete marks a Todo item as completed and generates the TodoCompleted event
func (t *Todo) Complete() error {
	if t.Status == "completed" {
		return errors.New("todo item is already completed")
	}
	completedAt := time.Now()
	t.Status = "completed"
	t.CompletedAt = &completedAt
	return nil
}

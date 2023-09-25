package domain

import "time"

// Event is the interface that all events must implement
type Event interface {
	AggregateID() string
	OccurredOn() time.Time
}

// TodoCreated is the event representing the creation of a Todo item
type TodoCreated struct {
	ID          string    // Unique identifier for the Todo item
	Title       string    // Title of the Todo item
	Description string    // Description of the Todo item
	Status      string    // Status of the Todo item
	CreatedAt   time.Time // Timestamp representing when the Todo item was created
}

// AggregateID returns the aggregate ID of the TodoCreated event
func (e *TodoCreated) AggregateID() string {
	return e.ID
}

// OccurredOn returns the timestamp of when the TodoCreated event occurred
func (e *TodoCreated) OccurredOn() time.Time {
	return e.CreatedAt
}

// TodoCompleted is the event representing the completion of a Todo item
type TodoCompleted struct {
	ID        string    // Unique identifier for the Todo item
	CompletedAt time.Time // Timestamp representing when the Todo item was completed
}

// AggregateID returns the aggregate ID of the TodoCompleted event
func (e *TodoCompleted) AggregateID() string {
	return e.ID
}

// OccurredOn returns the timestamp of when the TodoCompleted event occurred
func (e *TodoCompleted) OccurredOn() time.Time {
	return e.CompletedAt
}

// Other domain events can be defined here, for example, TodoDeleted, TodoUpdated, etc.

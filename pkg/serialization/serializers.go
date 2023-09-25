package serialization

import (
	"encoding/json"
	"errors"

	"project_name/pkg/domain"
)

// EventSerializer is an interface defining the methods that an event serializer should have.
type EventSerializer interface {
	Serialize(event domain.Event) ([]byte, error)
	Deserialize(data []byte, eventType string) (domain.Event, error)
}

// JsonSerializer is a simple, JSON implementation of the EventSerializer interface.
type JsonSerializer struct{}

// NewJsonSerializer initializes a new JsonSerializer.
func NewJsonSerializer() *JsonSerializer {
	return &JsonSerializer{}
}

// Serialize converts a domain event to a JSON byte slice.
func (s *JsonSerializer) Serialize(event domain.Event) ([]byte, error) {
	return json.Marshal(event)
}

// Deserialize converts a JSON byte slice to a domain event based on the eventType.
func (s *JsonSerializer) Deserialize(data []byte, eventType string) (domain.Event, error) {
	var event domain.Event

	switch eventType {
	case "TodoCreated":
		event = &domain.TodoCreated{}
	case "TodoCompleted":
		event = &domain.TodoCompleted{}
	default:
		return nil, errors.New("unknown event type")
	}

	if err := json.Unmarshal(data, event); err != nil {
		return nil, err
	}

	return event, nil
}

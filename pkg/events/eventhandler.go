package events

import (
	"sync"
)

// TaskReadModel represents a read model for a task.
type TaskReadModel struct {
	ID        string
	Title     string
	Completed bool
}

// EventHandler handles events and updates the read models/projections.
type EventHandler struct {
	mu     sync.RWMutex
	tasks  map[string]*TaskReadModel // In-memory read model
}

// NewEventHandler initializes a new EventHandler.
func NewEventHandler() *EventHandler {
	return &EventHandler{
		tasks: make(map[string]*TaskReadModel),
	}
}

// HandleEvent handles the given event.
func (h *EventHandler) HandleEvent(event Event) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	switch e := event.(type) {
	case *TaskCreatedEvent:
		return h.handleTaskCreated(e)
	case *TaskCompletedEvent:
		return h.handleTaskCompleted(e)
	case *TaskDeletedEvent:
		return h.handleTaskDeleted(e)
	default:
		return ErrInvalidEventType
	}
}

// handleTaskCreated handles TaskCreatedEvent.
func (h *EventHandler) handleTaskCreated(event *TaskCreatedEvent) error {
	h.tasks[event.ID] = &TaskReadModel{
		ID:    event.ID,
		Title: event.Title,
	}
	return nil
}

// handleTaskCompleted handles TaskCompletedEvent.
func (h *EventHandler) handleTaskCompleted(event *TaskCompletedEvent) error {
	task, exists := h.tasks[event.ID]
	if !exists {
		return ErrEventNotFound
	}
	task.Completed = true
	return nil
}

// handleTaskDeleted handles TaskDeletedEvent.
func (h *EventHandler) handleTaskDeleted(event *TaskDeletedEvent) error {
	delete(h.tasks, event.ID)
	return nil
}

// GetTask returns the read model for a task with the given ID.
func (h *EventHandler) GetTask(id string) (*TaskReadModel, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	task, exists := h.tasks[id]
	if !exists {
		return nil, ErrEventNotFound
	}
	return task, nil
}

// GetAllTasks returns all task read models.
func (h *EventHandler) GetAllTasks() []*TaskReadModel {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var tasks []*TaskReadModel
	for _, task := range h.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

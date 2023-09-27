package events 

import (
	"sync"
)

// TaskProjection represents a read model for a task.
type TaskProjection struct {
	ID        string
	Title     string
	Completed bool
	Deleted   bool
}

// ProjectionManager manages the read models/projections.
type ProjectionManager struct {
	mu    sync.RWMutex
	tasks map[string]*TaskProjection // In-memory read model
}

// NewProjectionManager initializes a new ProjectionManager.
func NewProjectionManager() *ProjectionManager {
	return &ProjectionManager{
		tasks: make(map[string]*TaskProjection),
	}
}

// HandleEvent handles the given event and updates the projections accordingly.
func (pm *ProjectionManager) HandleEvent(event Event) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	switch e := event.(type) {
	case *TaskCreatedEvent:
		pm.tasks[e.ID] = &TaskProjection{
			ID:    e.ID,
			Title: e.Title,
		}
	case *TaskCompletedEvent:
		task, exists := pm.tasks[e.ID]
		if !exists {
			return ErrEventNotFound
		}
		task.Completed = true
	case *TaskDeletedEvent:
		task, exists := pm.tasks[e.ID]
		if !exists {
			return ErrEventNotFound
		}
		task.Deleted = true
	default:
		return ErrInvalidEventType
	}
	return nil
}

// GetTask returns the projection for a task with the given ID.
func (pm *ProjectionManager) GetTask(id string) (*TaskProjection, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	task, exists := pm.tasks[id]
	if !exists {
		return nil, ErrEventNotFound
	}
	return task, nil
}

// GetAllTasks returns all task projections.
func (pm *ProjectionManager) GetAllTasks() []*TaskProjection {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	var tasks []*TaskProjection
	for _, task := range pm.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

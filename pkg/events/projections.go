package events

import (
	"sync"
)


type TaskProjection struct {
	ID        string
	Title     string
	Completed bool
	Deleted   bool
}


type ProjectionManager struct {
	mu    sync.RWMutex
	tasks map[string]*TaskProjection // In-memory read model
}


func NewProjectionManager() *ProjectionManager {
	return &ProjectionManager{
		tasks: make(map[string]*TaskProjection),
	}
}


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


func (pm *ProjectionManager) GetTask(id string) (*TaskProjection, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	task, exists := pm.tasks[id]
	if !exists {
		return nil, ErrEventNotFound
	}
	return task, nil
}


func (pm *ProjectionManager) GetAllTasks() []*TaskProjection {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	var tasks []*TaskProjection
	for _, task := range pm.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

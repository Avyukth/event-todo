package db

import (
	"sync"

	"event-todo/pkg/events"
)

// InMemoryDB represents an in-memory database.
type InMemoryDB struct {
	mu     sync.RWMutex
	tasks  map[string]*events.TaskProjection
}

// NewInMemoryDB initializes a new InMemoryDB.
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		tasks: make(map[string]*events.TaskProjection),
	}
}

// SaveTask saves a task projection to the in-memory database.
func (db *InMemoryDB) SaveTask(task *events.TaskProjection) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.tasks[task.ID] = task
}

// GetTask retrieves a task projection from the in-memory database by ID.
func (db *InMemoryDB) GetTask(id string) (*events.TaskProjection, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	task, exists := db.tasks[id]
	if !exists {
		return nil, events.ErrEventNotFound
	}
	return task, nil
}

// GetAllTasks retrieves all task projections from the in-memory database.
func (db *InMemoryDB) GetAllTasks() []*events.TaskProjection {
	db.mu.RLock()
	defer db.mu.RUnlock()

	var tasks []*events.TaskProjection
	for _, task := range db.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// DeleteTask deletes a task projection from the in-memory database by ID.
func (db *InMemoryDB) DeleteTask(id string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.tasks, id)
}

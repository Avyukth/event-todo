package db

import (
	"sync"

	"event-todo/pkg/todo"
)

// InMemoryDB represents an in-memory database.
type InMemoryDB struct {
	mu     sync.RWMutex
	tasks  map[string]*todo.TaskProjection
}

// NewInMemoryDB initializes a new InMemoryDB.
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		tasks: make(map[string]*todo.TaskProjection),
	}
}

// SaveTask saves a task projection to the in-memory database.
func (db *InMemoryDB) SaveTask(task *todo.TaskProjection) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.tasks[task.ID] = task
}

// GetTask retrieves a task projection from the in-memory database by ID.
func (db *InMemoryDB) GetTask(id string) (*todo.TaskProjection, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	task, exists := db.tasks[id]
	if !exists {
		return nil, todo.ErrTaskNotFound
	}
	return task, nil
}

// GetAllTasks retrieves all task projections from the in-memory database.
func (db *InMemoryDB) GetAllTasks() []*todo.TaskProjection {
	db.mu.RLock()
	defer db.mu.RUnlock()

	var tasks []*todo.TaskProjection
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

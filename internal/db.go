package db

import (
	"sync"

	"event-todo/pkg/events"
)


type InMemoryDB struct {
	mu     sync.RWMutex
	tasks  map[string]*events.TaskProjection
}


func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		tasks: make(map[string]*events.TaskProjection),
	}
}


func (db *InMemoryDB) SaveTask(task *events.TaskProjection) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.tasks[task.ID] = task
}


func (db *InMemoryDB) GetTask(id string) (*events.TaskProjection, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	task, exists := db.tasks[id]
	if !exists {
		return nil, events.ErrEventNotFound
	}
	return task, nil
}


func (db *InMemoryDB) GetAllTasks() []*events.TaskProjection {
	db.mu.RLock()
	defer db.mu.RUnlock()

	var tasks []*events.TaskProjection
	for _, task := range db.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}


func (db *InMemoryDB) DeleteTask(id string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.tasks, id)
}

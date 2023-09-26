package repository

import (
	"errors"
	"sync"

	domain "event-todo/pkg/domains"
)

// ReadModelRepository is an interface defining the methods that a read model repository should have.
type ReadModelRepository interface {
	GetTodoByID(id string) (*domain.TodoReadModel, error)
	SaveTodo(todo *domain.TodoReadModel) error
	GetAllTodos() ([]*domain.TodoReadModel, error)
}

// InMemoryReadModelRepository is a simple, in-memory implementation of the ReadModelRepository interface.
type InMemoryReadModelRepository struct {
	mu      sync.RWMutex
	storage map[string]*domain.TodoReadModel // map[id]TodoReadModel
}

// NewInMemoryReadModelRepository initializes a new InMemoryReadModelRepository.
func NewRepository() *InMemoryReadModelRepository {
	return &InMemoryReadModelRepository{
		storage: make(map[string]*domain.TodoReadModel),
	}
}

// GetTodoByID retrieves a TodoReadModel by its ID.
func (r *InMemoryReadModelRepository) GetTodoByID(id string) (*domain.TodoReadModel, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if todo, ok := r.storage[id]; ok {
		return todo, nil
	}

	return nil, errors.New("no todo found for the given ID")
}

// SaveTodo stores the provided TodoReadModel.
func (r *InMemoryReadModelRepository) SaveTodo(todo *domain.TodoReadModel) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.storage[todo.ID] = todo
	return nil
}

// GetAllTodos retrieves all TodoReadModels.
func (r *InMemoryReadModelRepository) GetAllTodos() ([]*domain.TodoReadModel, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var todos []*domain.TodoReadModel
	for _, todo := range r.storage {
		todos = append(todos, todo)
	}

	return todos, nil
}

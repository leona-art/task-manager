package inmemory

import (
	"fmt"

	"github.com/leona-art/task-manager/internal/domain/entity/task"
	"github.com/leona-art/task-manager/internal/domain/entity/todo"
)

type InMemoryTodoRepository struct {
	todos map[task.TaskId]todo.TodoTask
}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos: make(map[task.TaskId]todo.TodoTask),
	}
}

// Delete implements todo.TodoRepository.
func (i *InMemoryTodoRepository) Delete(id task.TaskId) (ok bool, err error) {
	delete(i.todos, id)
	return true, nil
}

// GetByID implements todo.TodoRepository.
func (i *InMemoryTodoRepository) Get(id task.TaskId) (value todo.TodoTask, ok bool, err error) {
	value, ok = i.todos[id]
	return value, ok, nil
}

// List implements todo.TodoRepository.
func (i *InMemoryTodoRepository) List() (todos []todo.TodoTask, err error) {
	for _, todo := range i.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

// Save implements todo.TodoRepository.
func (i *InMemoryTodoRepository) Save(todo todo.TodoTask) error {
	i.todos[todo.Data().ID] = todo
	return nil
}

func (i *InMemoryTodoRepository) Create(todo todo.TodoTask) error {
	if _, exists := i.todos[todo.Data().ID]; exists {
		return fmt.Errorf("todo with ID %s already exists", todo.Data().ID)
	}
	i.todos[todo.Data().ID] = todo
	return nil
}

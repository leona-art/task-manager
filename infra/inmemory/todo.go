package inmemory

import "github.com/leona-art/task-manager/domain/entity/todo"

type InMemoryTodoRepository struct {
	todos map[todo.TodoId]todo.Todo
}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{
		todos: make(map[todo.TodoId]todo.Todo),
	}
}

// Delete implements todo.TodoRepository.
func (i *InMemoryTodoRepository) Delete(id todo.TodoId) (ok bool, err error) {
	delete(i.todos, id)
	return true, nil
}

// GetByID implements todo.TodoRepository.
func (i *InMemoryTodoRepository) GetByID(id todo.TodoId) (value todo.Todo, ok bool, err error) {
	value, ok = i.todos[id]
	return value, ok, nil
}

// List implements todo.TodoRepository.
func (i *InMemoryTodoRepository) List() (todos []todo.Todo, err error) {
	for _, todo := range i.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

// Save implements todo.TodoRepository.
func (i *InMemoryTodoRepository) Save(todo todo.Todo) error {
	i.todos[todo.ID] = todo
	return nil
}

// Update implements todo.TodoRepository.
func (i *InMemoryTodoRepository) Update(todo todo.Todo) error {
	i.todos[todo.ID] = todo
	return nil
}

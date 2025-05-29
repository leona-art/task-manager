package usecase

import (
	"fmt"

	"github.com/leona-art/task-manager/domain/entity/todo"
)

type TodoUseCase struct {
	repository todo.TodoRepository
}

func NewTodoUseCase(repository todo.TodoRepository) *TodoUseCase {
	return &TodoUseCase{
		repository: repository,
	}
}

func (uc *TodoUseCase) CreateTodo(title, description string) (todo.Todo, error) {
	t, err := todo.NewTodo(title, description)
	if err != nil {
		return todo.Todo{}, err
	}
	if err := uc.repository.Save(t); err != nil {
		return todo.Todo{}, err
	}
	return t, nil
}

func (uc *TodoUseCase) GetTodoByID(id todo.TodoId) (todo.Todo, bool, error) {
	t, ok, err := uc.repository.GetByID(id)
	if err != nil {
		return todo.Todo{}, false, err
	}
	if !ok {
		return todo.Todo{}, false, nil
	}
	return t, true, nil
}

func (uc *TodoUseCase) UpdateTodo(t todo.Todo) error {
	if err := uc.repository.Update(t); err != nil {
		return err
	}
	return nil
}

func (uc *TodoUseCase) DeleteTodo(id todo.TodoId) (bool, error) {
	ok, err := uc.repository.Delete(id)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (uc *TodoUseCase) ListTodos() ([]todo.Todo, error) {
	todos, err := uc.repository.List()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (uc *TodoUseCase) DoTodo(id todo.TodoId) (todo.Todo, error) {
	t, ok, err := uc.GetTodoByID(id)
	if err != nil {
		return todo.Todo{}, err
	}
	if !ok {
		return todo.Todo{}, fmt.Errorf("todo with id %s not found", id)
	}
	if t.IsDone() {
		return todo.Todo{}, fmt.Errorf("todo with id %s is already done", id)
	}
	t.SwitchState()
	if err := uc.UpdateTodo(t); err != nil {
		return todo.Todo{}, err
	}
	return t, nil
}

func (uc *TodoUseCase) UndoTodo(id todo.TodoId) (todo.Todo, error) {
	t, ok, err := uc.GetTodoByID(id)
	if err != nil {
		return todo.Todo{}, err
	}
	if !ok {
		return todo.Todo{}, fmt.Errorf("todo with id %s not found", id)
	}
	if t.IsPending() {
		return todo.Todo{}, fmt.Errorf("todo with id %s is already pending", id)
	}
	t.SwitchState()
	if err := uc.UpdateTodo(t); err != nil {
		return todo.Todo{}, err
	}
	return t, nil
}

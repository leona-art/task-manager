package usecase

import (
	"fmt"

	"github.com/leona-art/task-manager/domain/entity/task"
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

func (uc *TodoUseCase) CreateTodo(title, description string) (todo.TodoTask, error) {
	taskEntity, err := task.NewTaskEntity(title, description)
	if err != nil {
		return todo.TodoTask{}, err
	}
	todoTask := todo.NewTodoTask(taskEntity)

	if err := uc.repository.Save(todoTask); err != nil {
		return todo.TodoTask{}, err
	}
	return todoTask, nil
}

func (uc *TodoUseCase) GetTodoByID(id task.TaskId) (todo.TodoTask, bool, error) {
	t, ok, err := uc.repository.Get(id)
	if err != nil {
		return todo.TodoTask{}, false, err
	}
	if !ok {
		return todo.TodoTask{}, false, nil
	}
	return t, true, nil
}

func (uc *TodoUseCase) UpdateTodo(t todo.TodoTask) error {
	if err := uc.repository.Save(t); err != nil {
		return err
	}
	return nil
}

func (uc *TodoUseCase) DeleteTodo(id task.TaskId) (bool, error) {
	ok, err := uc.repository.Delete(id)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (uc *TodoUseCase) ListTodos() ([]todo.TodoTask, error) {
	todos, err := uc.repository.List()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (uc *TodoUseCase) MarkAsDoneTodo(id task.TaskId) (todo.TodoTask, error) {
	t, ok, err := uc.GetTodoByID(id)
	if err != nil {
		return todo.TodoTask{}, err
	}
	if !ok {
		return todo.TodoTask{}, fmt.Errorf("todo with id %s not found", id)
	}
	if err := t.MarkAsDone(); err != nil {
		return todo.TodoTask{}, err
	}
	if err := uc.repository.Save(t); err != nil {
		return todo.TodoTask{}, err
	}
	return t, nil
}

func (uc *TodoUseCase) RevertTodo(id task.TaskId) (todo.TodoTask, error) {
	t, ok, err := uc.GetTodoByID(id)
	if err != nil {
		return todo.TodoTask{}, err
	}
	if !ok {
		return todo.TodoTask{}, fmt.Errorf("todo with id %s not found", id)
	}
	if err := t.Revert(); err != nil {
		return todo.TodoTask{}, err
	}
	if err := uc.repository.Save(t); err != nil {
		return todo.TodoTask{}, err
	}
	return t, nil
}

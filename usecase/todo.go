package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/leona-art/task-manager/adaptor/gateway"
	"github.com/leona-art/task-manager/domain/entity/task"
	"github.com/leona-art/task-manager/domain/entity/todo"
)

type TodoUseCase struct {
	repository gateway.TodoRepository
}

func NewTodoUseCase(repository gateway.TodoRepository) *TodoUseCase {
	return &TodoUseCase{
		repository: repository,
	}
}

func (uc *TodoUseCase) CreateTodo(ctx context.Context, title, description string) (todo.TodoTask, error) {
	taskEntity, err := task.NewTaskEntity(title, description)
	if err != nil {
		return todo.TodoTask{}, err
	}
	todoTask := todo.NewTodoTask(taskEntity)

	if err := uc.repository.Create(ctx, todoTask); err != nil {
		return todo.TodoTask{}, err
	}
	return todoTask, nil
}

func (uc *TodoUseCase) GetTodoByID(ctx context.Context, id string) (todo.TodoTask, bool, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return todo.TodoTask{}, false, err
	}
	t, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return todo.TodoTask{}, false, err
	}
	if !ok {
		return todo.TodoTask{}, false, nil
	}
	return t, true, nil
}

func (uc *TodoUseCase) UpdateTodo(ctx context.Context, t todo.TodoTask) error {
	if err := uc.repository.Save(ctx, t); err != nil {
		return err
	}
	return nil
}

func (uc *TodoUseCase) DeleteTodo(ctx context.Context, id string) (bool, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return false, err
	}
	ok, err := uc.repository.Delete(ctx, taskId)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (uc *TodoUseCase) ListTodos(ctx context.Context) ([]todo.TodoTask, error) {
	todos, err := uc.repository.List(ctx)
	if err != nil {
		slog.Error("failed to list todos", "error", err)
		return nil, err
	}
	return todos, nil
}

func (uc *TodoUseCase) MarkAsDoneTodo(ctx context.Context, id string) (todo.TodoTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return todo.TodoTask{}, err
	}
	t, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return todo.TodoTask{}, err
	}
	if !ok {
		return todo.TodoTask{}, fmt.Errorf("todo with id %s not found", id)
	}
	if err := t.MarkAsDone(); err != nil {
		return todo.TodoTask{}, err
	}
	if err := uc.repository.Save(ctx, t); err != nil {
		return todo.TodoTask{}, err
	}
	return t, nil
}

func (uc *TodoUseCase) RevertTodo(ctx context.Context, id string) (todo.TodoTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return todo.TodoTask{}, err
	}

	t, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return todo.TodoTask{}, err
	}
	if !ok {
		return todo.TodoTask{}, fmt.Errorf("todo with id %s not found", id)
	}
	if err := t.Revert(); err != nil {
		return todo.TodoTask{}, err
	}
	if err := uc.repository.Save(ctx, t); err != nil {
		return todo.TodoTask{}, err
	}
	return t, nil
}

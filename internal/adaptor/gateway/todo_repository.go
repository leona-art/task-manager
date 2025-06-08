package gateway

import (
	"context"

	"github.com/leona-art/task-manager/internal/domain/entity/task"
	"github.com/leona-art/task-manager/internal/domain/entity/todo"
)

type TodoRepository interface {
	Get(ctx context.Context, id task.TaskId) (task todo.TodoTask, ok bool, err error)
	Create(ctx context.Context, task todo.TodoTask) error
	Save(ctx context.Context, task todo.TodoTask) error
	Delete(ctx context.Context, id task.TaskId) (ok bool, err error)
	List(ctx context.Context) ([]todo.TodoTask, error)
}

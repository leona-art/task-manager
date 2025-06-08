package todo

import (
	"context"

	"github.com/leona-art/task-manager/domain/entity/task"
)

type TodoRepository interface {
	Get(ctx context.Context, id task.TaskId) (task TodoTask, ok bool, err error)
	Create(ctx context.Context, task TodoTask) error
	Save(ctx context.Context, task TodoTask) error
	Delete(ctx context.Context, id task.TaskId) (ok bool, err error)
	List(ctx context.Context) ([]TodoTask, error)
}

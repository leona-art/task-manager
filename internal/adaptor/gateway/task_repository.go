package gateway

import (
	"context"

	"github.com/leona-art/task-manager/internal/domain/entity/task"
)

type TaskRepository interface {
	Get(ctx context.Context, id task.TaskId) (task task.Task, ok bool, err error)
	Create(ctx context.Context, task task.Task) error
	Save(ctx context.Context, task task.Task) error
	Delete(ctx context.Context, id task.TaskId) (ok bool, err error)
	List(ctx context.Context) ([]task.Task, error)
}

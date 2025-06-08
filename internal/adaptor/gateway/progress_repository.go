package gateway

import (
	"context"

	"github.com/leona-art/task-manager/internal/domain/entity/progress"
	"github.com/leona-art/task-manager/internal/domain/entity/task"
)

type ProgressRepository interface {
	Get(ctx context.Context, id task.TaskId) (task progress.ProgressTask, ok bool, err error)
	Create(ctx context.Context, task progress.ProgressTask) error
	Save(ctx context.Context, task progress.ProgressTask) error
	Delete(ctx context.Context, id task.TaskId) (ok bool, err error)
	List(ctx context.Context) ([]progress.ProgressTask, error)
}

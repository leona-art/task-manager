package gateway

import (
	"github.com/leona-art/task-manager/internal/domain/entity/progress"
	"github.com/leona-art/task-manager/internal/domain/entity/task"
)

type ProgressRepository interface {
	Get(id task.TaskId) (task progress.ProgressTask, ok bool, err error)
	Create(task progress.ProgressTask) error
	Save(task progress.ProgressTask) error
	Delete(id task.TaskId) (ok bool, err error)
	List() ([]progress.ProgressTask, error)
}

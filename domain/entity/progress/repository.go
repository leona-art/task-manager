package progress

import "github.com/leona-art/task-manager/domain/entity/task"

type ProgressRepository interface {
	Get(id task.TaskId) (task ProgressTask, ok bool, err error)
	Create(task ProgressTask) error
	Save(task ProgressTask) error
	Delete(id task.TaskId) (ok bool, err error)
	List() ([]ProgressTask, error)
}

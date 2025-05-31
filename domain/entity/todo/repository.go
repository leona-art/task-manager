package todo

import "github.com/leona-art/task-manager/domain/entity/task"

type TodoRepository interface {
	Get(id task.TaskId) (task TodoTask, ok bool, err error)
	Create(task TodoTask) error
	Save(task TodoTask) error
	Delete(id task.TaskId) (ok bool, err error)
	List() ([]TodoTask, error)
}

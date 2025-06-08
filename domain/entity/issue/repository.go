package issue

import "github.com/leona-art/task-manager/domain/entity/task"

type IssueRepository interface {
	Get(id task.TaskId) (task IssueTask, ok bool, err error)
	Create(task IssueTask) error
	Save(task IssueTask) error
	Delete(id task.TaskId) (ok bool, err error)
	List() ([]IssueTask, error)
}

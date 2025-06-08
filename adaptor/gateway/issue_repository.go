package gateway

import (
	"github.com/leona-art/task-manager/domain/entity/issue"
	"github.com/leona-art/task-manager/domain/entity/task"
)

type IssueRepository interface {
	Get(id task.TaskId) (task issue.IssueTask, ok bool, err error)
	Create(task issue.IssueTask) error
	Save(task issue.IssueTask) error
	Delete(id task.TaskId) (ok bool, err error)
	List() ([]issue.IssueTask, error)
}

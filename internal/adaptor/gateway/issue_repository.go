package gateway

import (
	"context"

	"github.com/leona-art/task-manager/internal/domain/entity/issue"
	"github.com/leona-art/task-manager/internal/domain/entity/task"
)

type IssueRepository interface {
	Get(ctx context.Context, id task.TaskId) (task issue.IssueTask, ok bool, err error)
	Create(ctx context.Context, task issue.IssueTask) error
	Save(ctx context.Context, task issue.IssueTask) error
	Delete(ctx context.Context, id task.TaskId) (ok bool, err error)
	List(ctx context.Context) ([]issue.IssueTask, error)
}

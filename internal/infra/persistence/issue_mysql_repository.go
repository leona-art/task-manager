package persistence

import (
	"context"

	"github.com/leona-art/task-manager/gen/infra/sqlc"
	"github.com/leona-art/task-manager/internal/domain/entity/issue"
	"github.com/leona-art/task-manager/internal/domain/entity/task"
)

type IssueMySQLRepository struct {
	queries *sqlc.Queries
}

// Create implements gateway.IssueRepository.
func (i *IssueMySQLRepository) Create(ctx context.Context, task issue.IssueTask) error {
	panic("unimplemented")
}

// Delete implements gateway.IssueRepository.
func (i *IssueMySQLRepository) Delete(ctx context.Context, id task.TaskId) (ok bool, err error) {
	panic("unimplemented")
}

// Get implements gateway.IssueRepository.
func (i *IssueMySQLRepository) Get(ctx context.Context, id task.TaskId) (task issue.IssueTask, ok bool, err error) {
	panic("unimplemented")
}

// List implements gateway.IssueRepository.
func (i *IssueMySQLRepository) List(ctx context.Context) ([]issue.IssueTask, error) {
	panic("unimplemented")
}

// Save implements gateway.IssueRepository.
func (i *IssueMySQLRepository) Save(ctx context.Context, task issue.IssueTask) error {
	panic("unimplemented")
}

func NewMySqlIssueRepository(conn sqlc.DBTX) *IssueMySQLRepository {
	return &IssueMySQLRepository{
		queries: sqlc.New(conn),
	}
}

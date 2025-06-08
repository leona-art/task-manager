package persistence

import "github.com/leona-art/task-manager/gen/infra/sqlc"

type TaskMySQLRepository struct {
	queries *sqlc.Queries
}

func NewTaskMySQLRepository(conn sqlc.DBTX) *TaskMySQLRepository {
	return &TaskMySQLRepository{
		queries: sqlc.New(conn),
	}
}

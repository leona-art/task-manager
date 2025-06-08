package persistence

import (
	"context"
	"database/sql"

	"github.com/leona-art/task-manager/gen/infra/sqlc"
	"github.com/leona-art/task-manager/internal/domain/entity/progress"
	"github.com/leona-art/task-manager/internal/domain/entity/task"
	"github.com/leona-art/task-manager/internal/utils"
)

type ProgressMySQLRepository struct {
	queries *sqlc.Queries
}

// Create implements gateway.ProgressRepository.
func (p *ProgressMySQLRepository) Create(ctx context.Context, task progress.ProgressTask) error {

	err := p.queries.CreateTask(ctx, sqlc.CreateTaskParams{
		ID:          task.Data().ID.String(),
		Title:       task.Data().Title,
		Description: task.Data().Description,
		CreatedAt:   task.Data().CreatedAt,
		UpdatedAt:   task.Data().UpdatedAt,
		Kind:        sqlc.TasksKindProgress,
	})
	if err != nil {
		return err
	}
	err = p.queries.CreateProgressTask(ctx, sqlc.CreateProgressTaskParams{
		ID:       task.Data().ID.String(),
		Status:   sqlc.ProgressStatus(task.State().Status()),
		Solution: sql.NullString{},
	})
	if err != nil {
		return err
	}
	return nil
}

// Delete implements gateway.ProgressRepository.
func (p *ProgressMySQLRepository) Delete(ctx context.Context, id task.TaskId) (ok bool, err error) {
	err = p.queries.DeleteTask(ctx, id.String())
	if err != nil {
		return false, err
	}
	return true, nil
}

// Get implements gateway.ProgressRepository.
func (p *ProgressMySQLRepository) Get(ctx context.Context, id task.TaskId) (task progress.ProgressTask, ok bool, err error) {
	t, err := p.queries.GetProgressTask(ctx, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return progress.ProgressTask{}, false, nil // Task not found
		}
		return progress.ProgressTask{}, false, err // Other error
	}

	var solution utils.Option[string]
	if t.Solution.Valid {
		solution = utils.Some(t.Solution.String)
	} else {
		solution = utils.None[string]()
	}

	dto := progress.ProgressTaskDto{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      string(t.Status),
		Solution:    solution,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
	res, err := progress.NewProgressTaskFromDto(dto)
	if err != nil {
		return progress.ProgressTask{}, false, err // Invalid task data
	}
	return *res, true, nil // Task found
}

// List implements gateway.ProgressRepository.
func (p *ProgressMySQLRepository) List(ctx context.Context) ([]progress.ProgressTask, error) {
	progressTasks, err := p.queries.ListProgressTasks(ctx)
	if err != nil {
		return nil, err
	}
	tasks := make([]progress.ProgressTask, 0, len(progressTasks))
	for _, t := range progressTasks {
		var solution utils.Option[string]
		if t.Solution.Valid {
			solution = utils.Some(t.Solution.String)
		} else {
			solution = utils.None[string]()
		}
		dto := progress.ProgressTaskDto{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      string(t.Status),
			Solution:    solution,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		}
		res, err := progress.NewProgressTaskFromDto(dto)
		if err != nil {
			return nil, err // Invalid task data
		}
		tasks = append(tasks, *res)
	}
	return tasks, nil
}

// Save implements gateway.ProgressRepository.
func (p *ProgressMySQLRepository) Save(ctx context.Context, task progress.ProgressTask) error {
	err := p.queries.UpdateTask(ctx, sqlc.UpdateTaskParams{
		ID:          task.Data().ID.String(),
		Title:       task.Data().Title,
		Description: task.Data().Description,
		Kind:        sqlc.TasksKindProgress,
		UpdatedAt:   task.Data().UpdatedAt,
	})
	if err != nil {
		return err
	}
	solution, ok := task.State().Solution()
	err = p.queries.UpdateProgressStatus(ctx, sqlc.UpdateProgressStatusParams{
		ID:     task.Data().ID.String(),
		Status: sqlc.ProgressStatus(task.State().Status()),
		Solution: sql.NullString{
			String: solution,
			Valid:  ok,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func NewMySqlProgressRepository(conn sqlc.DBTX) *ProgressMySQLRepository {
	return &ProgressMySQLRepository{
		queries: sqlc.New(conn),
	}
}

package persistence

import (
	"context"
	"fmt"

	"github.com/leona-art/task-manager/domain/entity/task"
	"github.com/leona-art/task-manager/domain/entity/todo"
	"github.com/leona-art/task-manager/gen/infra/sqlc"
)

type TodoMySqlRepository struct {
	queries *sqlc.Queries
}

func NewMySqlTodoRepository(conn sqlc.DBTX) *TodoMySqlRepository {
	return &TodoMySqlRepository{
		queries: sqlc.New(conn),
	}
}
func (r *TodoMySqlRepository) Create(ctx context.Context, t todo.TodoTask) error {

	err := r.queries.CreateTask(ctx, sqlc.CreateTaskParams{
		ID:          t.Data().ID.String(),
		Title:       t.Data().Title,
		Description: t.Data().Description,
		Kind:        sqlc.TasksKind(t.Kind()),
		CreatedAt:   t.Data().CreatedAt,
		UpdatedAt:   t.Data().UpdatedAt,
	})
	if err != nil {
		return err
	}
	var status sqlc.TodosStatus
	switch t.State().Status() {
	case todo.Pending:
		status = sqlc.TodosStatusPending
	case todo.Done:
		status = sqlc.TodosStatusDone
	default:
		return fmt.Errorf("invalid todo status: %s", t.State().Status())
	}
	err = r.queries.CreateTodoTask(ctx, sqlc.CreateTodoTaskParams{
		ID:     t.Data().ID.String(),
		Status: status,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoMySqlRepository) Get(ctx context.Context, id task.TaskId) (todo.TodoTask, bool, error) {
	task, err := r.queries.GetTodoTask(ctx, id.String())
	if err != nil {
		return todo.TodoTask{}, false, err
	}

	todoTask, err := todo.NewTodoTaskFromDto(todo.TodoTaskDto{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		Status:      string(task.Status),
	})
	if err != nil {
		return todo.TodoTask{}, false, err
	}

	return todoTask, true, nil
}

func (r *TodoMySqlRepository) Save(ctx context.Context, t todo.TodoTask) error {
	err := r.queries.UpdateTask(ctx, sqlc.UpdateTaskParams{
		ID:          t.Data().ID.String(),
		Title:       t.Data().Title,
		Description: t.Data().Description,
		Kind:        sqlc.TasksKind(t.Kind()),
		UpdatedAt:   t.Data().UpdatedAt,
	})
	if err != nil {
		return err
	}
	var status sqlc.TodosStatus
	switch t.State().Status() {
	case todo.Pending:
		status = sqlc.TodosStatusPending
	case todo.Done:
		status = sqlc.TodosStatusDone
	default:
		return fmt.Errorf("invalid todo status: %s", t.State().Status())
	}
	err = r.queries.UpdateTodoStatus(ctx, sqlc.UpdateTodoStatusParams{
		ID:     t.Data().ID.String(),
		Status: status,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoMySqlRepository) Delete(ctx context.Context, id task.TaskId) (bool, error) {
	err := r.queries.DeleteTodoTask(ctx, id.String())
	if err != nil {
		return false, err
	}
	err = r.queries.DeleteTask(ctx, id.String())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *TodoMySqlRepository) List(ctx context.Context) ([]todo.TodoTask, error) {
	tasks, err := r.queries.ListTodoTasks(ctx)
	if err != nil {
		return nil, err
	}

	var todoTasks []todo.TodoTask
	for _, task := range tasks {
		todoTask, err := todo.NewTodoTaskFromDto(todo.TodoTaskDto{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
			Status:      string(task.Status),
		})
		if err != nil {
			return nil, err
		}
		todoTasks = append(todoTasks, todoTask)
	}

	return todoTasks, nil
}

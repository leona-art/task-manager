package todo

import (
	"fmt"
	"time"

	"github.com/leona-art/task-manager/domain/entity/task"
)

type TodoTask struct {
	data  task.TaskEntity
	state TodoState
}

func NewTodoTask(data task.TaskEntity) TodoTask {
	return TodoTask{
		data:  data,
		state: NewPendingState(),
	}
}

type TodoDto struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      string    `json:"status"`
}

func NewTodoTaskFromDto(dto TodoDto) (TodoTask, error) {
	if dto.ID == "" {
		return TodoTask{}, fmt.Errorf("task ID cannot be empty")
	}
	if dto.Title == "" {
		return TodoTask{}, fmt.Errorf("task title cannot be empty")
	}
	if dto.CreatedAt.IsZero() {
		return TodoTask{}, fmt.Errorf("task created_at cannot be empty")
	}
	if dto.UpdatedAt.IsZero() {
		return TodoTask{}, fmt.Errorf("task updated_at cannot be empty")
	}
	if dto.Status == "" {
		return TodoTask{}, fmt.Errorf("task status cannot be empty")
	}
	id, err := task.NewTaskIdFromString(dto.ID)
	if err != nil {
		return TodoTask{}, fmt.Errorf("invalid task ID: %v", err)
	}
	var state TodoState
	switch dto.Status {
	case "pending":
		state = NewPendingState()
	case "done":
		state = NewDoneState()
	default:
		return TodoTask{}, fmt.Errorf("invalid task status: %s", dto.Status)
	}
	return TodoTask{
		data: task.TaskEntity{
			ID:          id,
			Title:       dto.Title,
			Description: dto.Description,
			CreatedAt:   dto.CreatedAt,
			UpdatedAt:   dto.UpdatedAt,
		},
		state: state,
	}, nil
}

func (t *TodoTask) Data() task.TaskEntity {
	return t.data
}
func (t *TodoTask) Kind() task.TaskKind {
	return task.TaskKindTodo
}
func (t *TodoTask) State() TodoState {
	return t.state
}

func (t *TodoTask) MarkAsDone() error {
	if next, ok := t.state.Candidates()[Done]; ok {
		t.state = next()
	} else {
		return fmt.Errorf("cannot mark todo as done")
	}
	t.data.Update()
	return nil
}

func (t *TodoTask) Revert() error {
	if next, ok := t.state.Candidates()[Pending]; ok {
		t.state = next()
	} else {
		return fmt.Errorf("cannot pend todo")
	}
	t.data.Update()
	return nil
}

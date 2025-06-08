package todo

import (
	"fmt"

	"github.com/leona-art/task-manager/internal/domain/entity/task"
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

func NewTodoTaskFromDto(dto TodoTaskDto) (TodoTask, error) {
	id, err := task.NewTaskIdFromString(dto.ID)
	if err != nil {
		return TodoTask{}, fmt.Errorf("invalid task ID: %v", err)
	}
	data := task.TaskEntity{
		ID:          id,
		Title:       dto.Title,
		Description: dto.Description,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
	if !data.IsValid() {
		return TodoTask{}, fmt.Errorf("invalid task data: %v", data)
	}
	if data.IsEmpty() {
		return TodoTask{}, fmt.Errorf("task data cannot be empty")
	}
	state := NewTodoState(dto.Status)
	if state == nil {
		return TodoTask{}, fmt.Errorf("invalid task status: %s", dto.Status)
	}
	return TodoTask{
		data:  data,
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

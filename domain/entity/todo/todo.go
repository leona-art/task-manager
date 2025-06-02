package todo

import (
	"fmt"

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

func (t *TodoTask) Data() task.TaskEntity {
	return t.data
}
func (t *TodoTask) Kind() task.TaskKind {
	return task.TaskKindTodo
}

func (t *TodoTask) MarkAsDone() error {
	if next, ok := t.state.Candidates()[Done]; ok {
		t.state = next()
	} else {
		return fmt.Errorf("cannot mark todo as done")
	}
	return nil
}

func (t *TodoTask) Revert() error {
	if next, ok := t.state.Candidates()[Pending]; ok {
		t.state = next()
	} else {
		return fmt.Errorf("cannot pend todo")
	}
	return nil
}

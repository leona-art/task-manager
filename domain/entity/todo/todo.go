package todo

import (
	"fmt"

	"github.com/leona-art/task-manager/domain/entity/task"
)

type TodoTask struct {
	info  task.BaseTask
	state TodoState
}

func (t *TodoTask) Info() task.BaseTask {
	return t.info
}
func (t *TodoTask) Kind() task.TaskKind {
	return task.TaskKindTodo
}

func (t *TodoTask) Do() error {
	if next, ok := t.state.Candidates()[Done]; ok {
		t.state = next()
	} else {
		return fmt.Errorf("cannot mark todo as done")
	}
	return nil
}

func (t *TodoTask) Pend() error {
	if next, ok := t.state.Candidates()[Pending]; ok {
		t.state = next()
	} else {
		return fmt.Errorf("cannot pend todo")
	}
	return nil
}

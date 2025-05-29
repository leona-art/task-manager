package task

import (
	"fmt"

	"github.com/leona-art/task-manager/domain/entity/task_info"
)

type Task struct {
	Info  task_info.TaskInfo
	State TaskState
}

func NewTask(title, description string) (Task, error) {
	info, err := task_info.NewTaskInfo(title, description)
	if err != nil {
		return Task{}, err
	}
	return Task{
		Info:  info,
		State: NewTaskPendingState(),
	}, nil
}

func (t *Task) Start() error {
	candidate := t.State.Candidate()
	if progress, ok := candidate[InProgress]; ok {
		t.State = progress()
	} else {
		return fmt.Errorf("cannot start task: no candidate status found")
	}
	return nil
}

func (t *Task) SetResolution(resolution string) error {
	switch t.State.Status() {
	case InProgress:
		t.State = NewTaskInProgressStateWithResolution(resolution)
	default:
		return fmt.Errorf("cannot set resolution: task is not in progress")
	}
	return nil
}

func (t *Task) Complete(resolution string) error {
	candidate := t.State.Candidate()
	if complete, ok := candidate[Completed]; ok {
		t.State = complete()
	} else {
		return fmt.Errorf("cannot complete task: no candidate status found")
	}
	return nil
}

func (t *Task) Equal(other Task) bool {
	return t.Info.Equal(other.Info) &&
		t.State.Status() == other.State.Status()
}

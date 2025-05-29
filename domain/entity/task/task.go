package task

import "fmt"

type Task struct {
	ID          string
	Title       string
	Description string
	State       TaskState
}

func NewTask(id, title, description string) Task {
	return Task{
		ID:          id,
		Title:       title,
		Description: description,
		State:       NewTaskPendingState(),
	}
}

func (t *Task) Open() error {
	candidate := t.State.Candidate()
	if progress, ok := candidate[InProgress]; ok {
		t.State = progress()
	} else {
		return fmt.Errorf("cannot open task: no candidate status found")
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
	return t.ID == other.ID &&
		t.Title == other.Title &&
		t.Description == other.Description &&
		t.State.Status() == other.State.Status()
}

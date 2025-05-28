package task

import "fmt"

type Task struct {
	ID          string
	Title       string
	Description string
	Status      TaskStatus
}

func NewTask(id, title, description string) Task {
	return Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      NewTaskPendingStatus(),
	}
}

func (t *Task) Open() error {
	candidate := t.Status.Candidate()
	if progress, ok := candidate[InProgress]; ok {
		t.Status = progress()
	} else {
		return fmt.Errorf("cannot open task: no candidate status found")
	}
	return nil
}

func (t *Task) SetResolution(resolution string) error {
	switch t.Status.Status() {
	case InProgress:
		t.Status = NewTaskInProgressStatusWithResolution(resolution)
	default:
		return fmt.Errorf("cannot set resolution: task is not in progress")
	}
	return nil
}

func (t *Task) Complete(resolution string) error {
	candidate := t.Status.Candidate()
	if complete, ok := candidate[Completed]; ok {
		t.Status = complete()
	} else {
		return fmt.Errorf("cannot complete task: no candidate status found")
	}
	return nil
}

func (t *Task) Equal(other Task) bool {
	return t.ID == other.ID &&
		t.Title == other.Title &&
		t.Description == other.Description &&
		t.Status.Status() == other.Status.Status()
}

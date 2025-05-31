package task

import "time"

type BaseTask struct {
	ID          TaskId
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewBaseTask(title, description string) (BaseTask, error) {
	id, err := NewTaskId()
	if err != nil {
		return BaseTask{}, err
	}
	now := time.Now()
	return BaseTask{
		ID:          id,
		Title:       title,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

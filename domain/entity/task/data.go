package task

import "time"

type TaskEntity struct {
	ID          TaskId
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTaskEntity(title, description string) (TaskEntity, error) {
	id, err := NewTaskId()
	if err != nil {
		return TaskEntity{}, err
	}
	now := time.Now()
	return TaskEntity{
		ID:          id,
		Title:       title,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

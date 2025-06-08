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

func (t *TaskEntity) Set(title, description string) {
	t.Title = title
	t.Description = description
	t.UpdatedAt = time.Now()
}

func (t *TaskEntity) Update() {
	t.UpdatedAt = time.Now()
}

func (t *TaskEntity) IsEmpty() bool {
	return t.ID.IsEmpty() && t.Title == "" && t.Description == "" && t.CreatedAt.IsZero() && t.UpdatedAt.IsZero()
}

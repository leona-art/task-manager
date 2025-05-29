package taskinfo

import (
	"time"
)

type TaskInfo struct {
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTaskInfo(title, description string) TaskInfo {
	return TaskInfo{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *TaskInfo) Set(title, description string) {
	t.Title = title
	t.Description = description
	t.UpdatedAt = time.Now()
}

func (t *TaskInfo) Update() {
	t.UpdatedAt = time.Now()
}

func (t *TaskInfo) Equal(other TaskInfo) bool {
	return t.Title == other.Title &&
		t.Description == other.Description &&
		t.CreatedAt.Equal(other.CreatedAt) &&
		t.UpdatedAt.Equal(other.UpdatedAt)
}

func (t *TaskInfo) IsEmpty() bool {
	return t.Title == "" && t.Description == "" &&
		t.CreatedAt.IsZero() && t.UpdatedAt.IsZero()
}

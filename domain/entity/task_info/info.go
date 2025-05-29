package task_info

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TaskInfo struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTaskInfo(title, description string) (TaskInfo, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return TaskInfo{}, fmt.Errorf("failed to generate UUID: %w", err)
	}
	return TaskInfo{
		ID:          id.String(),
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
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
	return t.ID == other.ID &&
		t.Title == other.Title &&
		t.Description == other.Description &&
		t.CreatedAt.Equal(other.CreatedAt) &&
		t.UpdatedAt.Equal(other.UpdatedAt)
}

func (t *TaskInfo) IsEmpty() bool {
	return t.ID == "" && t.Title == "" && t.Description == "" &&
		t.CreatedAt.IsZero() && t.UpdatedAt.IsZero()
}

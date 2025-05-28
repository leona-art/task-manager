package task_info

import (
	"time"

	"github.com/google/uuid"
)

type TaskInfo struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
}

func NewTaskInfo(title, description string) TaskInfo {
	id, err := uuid.NewV7()
	if err != nil {
		panic("failed to generate UUID: " + err.Error())
	}
	return TaskInfo{
		ID:          id.String(),
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}
}

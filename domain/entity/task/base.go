package task

import "time"

type BaseTask struct {
	ID          TaskId
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

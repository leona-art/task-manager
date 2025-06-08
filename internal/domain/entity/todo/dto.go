package todo

import "time"

type TodoTaskDto struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      string
}

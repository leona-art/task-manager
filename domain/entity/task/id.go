package task

import (
	"fmt"

	"github.com/google/uuid"
)

type TaskId string

func NewTaskId() (TaskId, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return TaskId(id.String()), nil
}

func NewTaskIdFromString(s string) (TaskId, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return "", err
	}
	if id.Version() != 7 {
		return "", fmt.Errorf("invalid UUID version: %d, expected version 7", id.Version())
	}
	return TaskId(s), nil
}

func (id TaskId) String() string {
	return string(id)
}
func (id TaskId) IsEmpty() bool {
	return id == ""
}
func (id TaskId) Equal(other TaskId) bool {
	return id == other
}

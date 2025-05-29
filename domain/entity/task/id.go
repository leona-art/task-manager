package task

import "github.com/google/uuid"

type TaskId string

func NewTaskId() (TaskId, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return TaskId(id.String()), nil
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

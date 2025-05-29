package todo

import "github.com/google/uuid"

type TodoId string

func (id TodoId) String() string {
	return string(id)
}
func NewTodoId() (TodoId, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return TodoId(id.String()), nil
}

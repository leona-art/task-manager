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

func (id TodoId) Equal(other TodoId) bool {
	return id.String() == other.String()
}
func (id TodoId) IsEmpty() bool {
	return id.String() == ""
}
func (id TodoId) IsValid() bool {
	_, err := uuid.Parse(id.String())
	return err == nil
}
func (id TodoId) IsZero() bool {
	return id.String() == ""
}

func NewTodoIdFromString(id string) (TodoId, error) {
	if id == "" {
		return "", nil
	}
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}
	return TodoId(parsedId.String()), nil
}

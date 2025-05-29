package todo

import "github.com/leona-art/task-manager/domain/entity/taskinfo"

type Todo struct {
	ID    TodoId
	Info  taskinfo.TaskInfo
	State TodoState
}

func NewTodo(title, description string) (Todo, error) {
	id, err := NewTodoId()
	if err != nil {
		return Todo{}, err
	}
	return Todo{
		ID:    id,
		Info:  taskinfo.NewTaskInfo(title, description),
		State: NewTodoPendingState(),
	}, nil
}
func (t *Todo) SwitchState() {
	t.State = t.State.Switch()
	t.Info.Update()
}

func (t *Todo) Equal(other Todo) bool {
	return t.Info.Equal(other.Info) &&
		t.State.Status() == other.State.Status()
}

func (t *Todo) IsDone() bool {
	return t.State.Status() == Done
}
func (t *Todo) IsPending() bool {
	return t.State.Status() == Pending
}

func (t *Todo) IsEmpty() bool {
	return t.ID.String() == "" &&
		t.Info.IsEmpty() &&
		t.State.Status() == Pending
}

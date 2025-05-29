package todo

import "github.com/leona-art/task-manager/domain/entity/taskinfo"

type Todo struct {
	Info  taskinfo.TaskInfo
	State TodoState
}

func NewTodo(title, description string) (Todo, error) {
	info, err := taskinfo.NewTaskInfo(title, description)
	if err != nil {
		return Todo{}, err
	}
	return Todo{
		Info:  info,
		State: NewTodoPendingState(),
	}, nil
}
func (t *Todo) SwitchState() {
	t.State = t.State.Switch()
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

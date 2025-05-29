package todo

import "github.com/leona-art/task-manager/domain/entity/task_info"

type Todo struct {
	Info  task_info.TaskInfo
	State TodoState
}

func NewTodo(title, description string) (Todo, error) {
	info, err := task_info.NewTaskInfo(title, description)
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

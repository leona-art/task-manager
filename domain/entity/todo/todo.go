package todo

type Todo struct {
	ID          string
	Title       string
	Description string
	State       TodoState
}

func NewTodo(id, title, description string) Todo {
	return Todo{
		ID:          id,
		Title:       title,
		Description: description,
		State:       NewTodoPendingState(),
	}
}
func (t *Todo) SwitchState() {
	t.State = t.State.Switch()
}

func (t *Todo) Equal(other Todo) bool {
	return t.ID == other.ID &&
		t.Title == other.Title &&
		t.Description == other.Description &&
		t.State.Status() == other.State.Status()
}

func (t *Todo) IsDone() bool {
	return t.State.Status() == Done
}
func (t *Todo) IsPending() bool {
	return t.State.Status() == Pending
}

package todo

type Todo struct {
	ID          string
	Title       string
	Description string
	Status      TodoStatus
}

func NewTodo(id, title, description string) Todo {
	return Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      NewTodoPendingStatus(),
	}
}
func (t *Todo) SwitchStatus() {
	t.Status = t.Status.Switch()
}

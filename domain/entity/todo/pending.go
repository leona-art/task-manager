package todo

type TodoPendingStatus struct{}

func (s TodoPendingStatus) Status() string {
	return "pending"
}

func (s TodoPendingStatus) Switch() TodoStatus {
	return NewTodoDoneStatus()
}

func NewTodoPendingStatus() TodoPendingStatus {
	return TodoPendingStatus{}
}

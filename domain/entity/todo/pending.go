package todo

type TodoPendingStatus struct{}

func (s TodoPendingStatus) Status() TodoStatus {
	return Pending
}

func (s TodoPendingStatus) Switch() TodoState {
	return NewTodoDoneState()
}

func NewTodoPendingState() TodoPendingStatus {
	return TodoPendingStatus{}
}

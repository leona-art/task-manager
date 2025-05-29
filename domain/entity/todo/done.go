package todo

type TodoDoneStatus struct{}

func (s TodoDoneStatus) Status() TodoStatus {
	return Done
}
func (s TodoDoneStatus) Switch() TodoState {
	return NewTodoPendingState()
}

func NewTodoDoneState() TodoDoneStatus {
	return TodoDoneStatus{}
}

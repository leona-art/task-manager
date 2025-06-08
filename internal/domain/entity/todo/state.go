package todo

type TransitionMap map[TodoStatus]func() TodoState
type TodoState interface {
	Status() TodoStatus
	Candidates() TransitionMap
}

func NewTodoState(status string) TodoState {
	switch status {
	case string(Pending):
		return NewPendingState()
	case string(Done):
		return NewDoneState()
	default:
		return nil
	}
}

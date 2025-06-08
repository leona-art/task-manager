package todo

type PendingState struct{}

func (s PendingState) Status() TodoStatus {
	return Pending
}
func (s PendingState) Candidates() TransitionMap {
	return TransitionMap{
		Done: func() TodoState {
			return NewDoneState()
		},
	}
}

func NewPendingState() TodoState {
	return PendingState{}
}

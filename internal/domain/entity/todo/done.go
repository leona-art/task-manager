package todo

type DoneState struct{}

func (s DoneState) Status() TodoStatus {
	return Done
}
func (s DoneState) Candidates() TransitionMap {
	return TransitionMap{
		Pending: func() TodoState {
			return NewPendingState()
		},
	}
}
func NewDoneState() TodoState {
	return DoneState{}
}

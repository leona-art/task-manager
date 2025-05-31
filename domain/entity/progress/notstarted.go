package progress

type NotStartedState struct{}

func (s NotStartedState) Status() ProgressStatus {
	return NotStarted
}

func (s NotStartedState) Solution() (value string, ok bool) {
	return "", false
}
func (s NotStartedState) Candidates() TransitionMap {
	return TransitionMap{
		InProgress: func() ProgressState {
			return NewInProgressState()
		},
	}
}

func NewNotStartedState() ProgressState {
	return NotStartedState{}
}

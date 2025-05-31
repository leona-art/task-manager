package progress

type CompletedState struct {
	solution string
}

func (s CompletedState) Status() ProgressStatus {
	return Completed
}

func (s CompletedState) Solution() (value string, ok bool) {
	return s.solution, true
}
func (s CompletedState) Candidates() TransitionMap {
	return TransitionMap{
		InProgress: func() ProgressState {
			return NewInProgressStateWithSolution(s.solution)
		},
	}
}

func NewCompletedState(solution string) ProgressState {
	return CompletedState{
		solution: solution,
	}
}

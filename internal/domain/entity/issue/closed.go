package issue

type ClosedState struct {
	cause    string
	solution string
}

func (s ClosedState) Status() IssueStatus {
	return Closed
}

func (s ClosedState) Cause() (value string, ok bool) {
	return s.cause, true
}
func (s ClosedState) Solution() (value string, ok bool) {
	return s.solution, true
}
func (s ClosedState) Candidates() TransitionMap {
	var TransitionMap = make(TransitionMap)
	TransitionMap[Investigating] = func() IssueState {
		return NewInvestigatingStateWithCause(s.cause)
	}
	TransitionMap[Resolving] = func() IssueState {
		return NewResolvingStateWithSolution(s.cause, s.solution)
	}
	return TransitionMap
}

func NewClosedState(cause, solution string) IssueState {
	return ClosedState{
		cause:    cause,
		solution: solution,
	}
}

package issue

import "github.com/leona-art/task-manager/internal/utils"

type ResolvingState struct {
	cause    string
	solution utils.Option[string]
}

func (s ResolvingState) Status() IssueStatus {
	return Resolving
}

func (s ResolvingState) Cause() (value string, ok bool) {
	return s.cause, true
}
func (s ResolvingState) Solution() (value string, ok bool) {
	return s.solution.Get()
}
func (s ResolvingState) Candidates() TransitionMap {
	var TransitionMap = make(TransitionMap)
	if value, ok := s.solution.Get(); ok {
		TransitionMap[Closed] = func() IssueState {
			return NewClosedState(s.cause, value)
		}
	}
	return TransitionMap
}
func NewResolvingState(cause string) IssueState {
	return ResolvingState{
		cause:    cause,
		solution: utils.None[string](),
	}
}

func NewResolvingStateWithSolution(cause, solution string) IssueState {
	return ResolvingState{
		cause:    cause,
		solution: utils.Some(solution),
	}
}

func (s ResolvingState) WithSolution(solution string) IssueState {
	return ResolvingState{
		cause:    s.cause,
		solution: utils.Some(solution),
	}
}

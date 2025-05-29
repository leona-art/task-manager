package issue

import "github.com/leona-art/task-manager/utils"

type IssueResolvingStatus struct {
	cause      string
	resolution utils.Option[string]
}

func (s IssueResolvingStatus) Status() IssueStatus {
	return Resolving
}
func (s IssueResolvingStatus) Cause() (value string, ok bool) {
	return s.cause, true
}
func (s IssueResolvingStatus) Resolution() (value string, ok bool) {
	return s.resolution.Get()
}
func (s IssueResolvingStatus) Candidate() TransitionMap {
	var transitionMap = make(TransitionMap)
	transitionMap[Researching] = func() IssueState {
		return NewIssueResearchingStateWithCause(s.cause)
	}
	if value, ok := s.resolution.Get(); ok {
		transitionMap[Resolved] = func() IssueState {
			return NewIssueResolvedState(s.cause, value)
		}
	}
	return transitionMap
}
func NewIssueResolvingState(cause string) IssueResolvingStatus {
	return IssueResolvingStatus{
		cause:      cause,
		resolution: utils.None[string](),
	}
}

func NewIssueResolvingStateWithResolution(cause string, resolution string) IssueResolvingStatus {
	return IssueResolvingStatus{
		cause:      cause,
		resolution: utils.Some(resolution),
	}
}

package issue

import "github.com/leona-art/task-manager/utils"

type IssueResearchingStatus struct {
	cause utils.Option[string]
}

func (s IssueResearchingStatus) Status() IssueStatus {
	return Researching
}
func (s IssueResearchingStatus) Cause() (value string, ok bool) {
	return s.cause.Get()
}
func (s IssueResearchingStatus) Resolution() (value string, ok bool) {
	return "", false
}
func (s IssueResearchingStatus) Candidate() TransitionMap {
	var transitionMap = make(TransitionMap)
	if value, ok := s.cause.Get(); ok {
		transitionMap[Resolving] = func() IssueState {
			return NewIssueResolvingState(value)
		}
	}
	return transitionMap
}
func NewIssueResearchingState() IssueResearchingStatus {
	return IssueResearchingStatus{
		cause: utils.None[string](),
	}
}

func NewIssueResearchingStateWithCause(cause string) IssueResearchingStatus {
	return IssueResearchingStatus{
		cause: utils.Some(cause),
	}
}

package issue

type IssueResolvedStatus struct {
	resolution string
	cause      string
}

func (s IssueResolvedStatus) Status() IssueStatus {
	return Resolved
}

func (s IssueResolvedStatus) Cause() (value string, ok bool) {
	return s.cause, true
}
func (s IssueResolvedStatus) Resolution() (value string, ok bool) {
	return s.resolution, true
}
func (s IssueResolvedStatus) Candidate() TransitionMap {
	return TransitionMap{
		Researching: func() IssueState {
			return NewIssueResearchingStateWithCause(s.cause)
		},
		Resolving: func() IssueState {
			return NewIssueResolvingStateWithResolution(s.cause, s.resolution)
		},
	}
}
func NewIssueResolvedState(cause, resolution string) IssueResolvedStatus {
	return IssueResolvedStatus{
		cause:      cause,
		resolution: resolution,
	}
}

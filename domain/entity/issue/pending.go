package issue

type IssuePendingStatus struct{}

func (s IssuePendingStatus) Status() IssueStatus {
	return Pending
}

func (s IssuePendingStatus) Cause() (value string, ok bool) {
	return "", false
}
func (s IssuePendingStatus) Resolution() (value string, ok bool) {
	return "", false
}
func (s IssuePendingStatus) Candidate() TransitionMap {
	return TransitionMap{
		Researching: func() IssueState {
			return NewIssueResearchingState()
		},
	}
}
func NewIssuePendingState() IssuePendingStatus {
	return IssuePendingStatus{}
}

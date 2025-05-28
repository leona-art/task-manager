package issue

type IssuePendingStatus struct{}

func (s IssuePendingStatus) Status() string {
	return Pending
}

func (s IssuePendingStatus) Cause() (value string, ok bool) {
	return "", false
}
func (s IssuePendingStatus) Resolution() (value string, ok bool) {
	return "", false
}
func (s IssuePendingStatus) Candidate() map[string]func() IssueStatus {
	return map[string]func() IssueStatus{
		Researching: func() IssueStatus {
			return NewIssueResearchingStatus()
		},
	}
}
func NewIssuePendingStatus() IssuePendingStatus {
	return IssuePendingStatus{}
}

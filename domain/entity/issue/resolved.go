package issue

type IssueResolvedStatus struct {
	resolution string
	cause      string
}

func (s IssueResolvedStatus) Status() string {
	return Resolved
}

func (s IssueResolvedStatus) Cause() (value string, ok bool) {
	return s.cause, true
}
func (s IssueResolvedStatus) Resolution() (value string, ok bool) {
	return s.resolution, true
}
func (s IssueResolvedStatus) Candidate() map[string]func() IssueStatus {
	return map[string]func() IssueStatus{
		Researching: func() IssueStatus {
			return NewIssueResearchingStatusWithCause(s.cause)
		},
		Resolving: func() IssueStatus {
			return NewIssueResolvingStatusWithResolution(s.cause, s.resolution)
		},
	}
}
func NewIssueResolvedStatus(cause, resolution string) IssueResolvedStatus {
	return IssueResolvedStatus{
		cause:      cause,
		resolution: resolution,
	}
}

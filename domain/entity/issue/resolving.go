package issue

import "github.com/leona-art/task-manager/utils"

type IssueResolvingStatus struct {
	cause      string
	resolution utils.Option[string]
}

func (s IssueResolvingStatus) Status() string {
	return Resolving
}
func (s IssueResolvingStatus) Cause() (value string, ok bool) {
	return s.cause, true
}
func (s IssueResolvingStatus) Resolution() (value string, ok bool) {
	return s.resolution.Get()
}
func (s IssueResolvingStatus) Candidate() map[string]func() IssueStatus {
	resolution, ok := s.resolution.Get()
	if !ok {
		return map[string]func() IssueStatus{
			Researching: func() IssueStatus {
				return NewIssueResearchingStatusWithCause(s.cause)
			},
		}
	}
	return map[string]func() IssueStatus{
		Researching: func() IssueStatus {
			return NewIssueResearchingStatusWithCause(s.cause)
		},
		Resolved: func() IssueStatus {
			return NewIssueResolvedStatus(s.cause, resolution)
		},
	}
}
func NewIssueResolvingStatus(cause string) IssueResolvingStatus {
	return IssueResolvingStatus{
		cause:      cause,
		resolution: utils.None[string](),
	}
}

func NewIssueResolvingStatusWithResolution(cause string, resolution string) IssueResolvingStatus {
	return IssueResolvingStatus{
		cause:      cause,
		resolution: utils.Some(resolution),
	}
}

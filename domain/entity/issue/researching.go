package issue

import "github.com/leona-art/task-manager/utils"

type IssueResearchingStatus struct {
	cause utils.Option[string]
}

func (s IssueResearchingStatus) Status() string {
	return Researching
}
func (s IssueResearchingStatus) Cause() (value string, ok bool) {
	return s.cause.Get()
}
func (s IssueResearchingStatus) Resolution() (value string, ok bool) {
	return "", false
}
func (s IssueResearchingStatus) Candidate() map[string]func() IssueStatus {
	value, ok := s.cause.Get()
	if !ok {
		return map[string]func() IssueStatus{}
	}
	return map[string]func() IssueStatus{
		Resolving: func() IssueStatus {
			return NewIssueResolvingStatus(value)
		},
	}
}
func NewIssueResearchingStatus() IssueResearchingStatus {
	return IssueResearchingStatus{
		cause: utils.None[string](),
	}
}

func NewIssueResearchingStatusWithCause(cause string) IssueResearchingStatus {
	return IssueResearchingStatus{
		cause: utils.Some(cause),
	}
}

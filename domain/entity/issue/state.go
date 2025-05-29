package issue

type IssueState interface {
	Status() IssueStatus
	Cause() (value string, ok bool)
	Resolution() (value string, ok bool)
	Candidate() TransitionMap
}
type TransitionMap map[IssueStatus]func() IssueState

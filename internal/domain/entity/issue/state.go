package issue

type TransitionMap map[IssueStatus]func() IssueState
type IssueState interface {
	Status() IssueStatus
	Candidates() TransitionMap
	Cause() (value string, ok bool)
	Solution() (value string, ok bool)
}

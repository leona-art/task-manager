package issue

type OpenStatus struct{}

func (s OpenStatus) Status() IssueStatus {
	return Open
}

func (s OpenStatus) Cause() (value string, ok bool) {
	return "", false
}
func (s OpenStatus) Solution() (value string, ok bool) {
	return "", false
}

func (s OpenStatus) Candidates() TransitionMap {
	return TransitionMap{
		Investigating: func() IssueState {
			return NewInvestigatingState()
		},
	}
}

func NewOpenState() IssueState {
	return OpenStatus{}
}

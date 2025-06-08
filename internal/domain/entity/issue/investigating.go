package issue

import "github.com/leona-art/task-manager/utils"

type InvestigatingState struct {
	cause utils.Option[string]
}

func (s InvestigatingState) Status() IssueStatus {
	return Investigating
}

func (s InvestigatingState) Cause() (value string, ok bool) {
	return s.cause.Get()
}

func (s InvestigatingState) Solution() (value string, ok bool) {
	return "", false
}

func (s InvestigatingState) Candidates() TransitionMap {
	var TransitionMap = make(TransitionMap)
	if value, ok := s.cause.Get(); ok {
		TransitionMap[Resolving] = func() IssueState {
			return NewResolvingState(value)
		}
	}
	return TransitionMap
}

func NewInvestigatingState() IssueState {
	return InvestigatingState{
		cause: utils.None[string](),
	}
}

func NewInvestigatingStateWithCause(cause string) IssueState {
	return InvestigatingState{
		cause: utils.Some(cause),
	}
}
func (s InvestigatingState) WithCause(cause string) IssueState {
	return InvestigatingState{
		cause: utils.Some(cause),
	}
}

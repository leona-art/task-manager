package progress

import "github.com/leona-art/task-manager/internal/utils"

type InProgressState struct {
	solution utils.Option[string]
}

func (s InProgressState) Status() ProgressStatus {
	return InProgress
}
func (s InProgressState) Solution() (value string, ok bool) {
	return s.solution.Get()
}

func (s InProgressState) Candidates() TransitionMap {
	var TransitionMap = make(TransitionMap)

	if value, ok := s.solution.Get(); ok {
		TransitionMap[Completed] = func() ProgressState {
			return NewCompletedState(value)
		}
	}
	return TransitionMap
}

func NewInProgressState() ProgressState {
	return InProgressState{
		solution: utils.None[string](),
	}
}

func NewInProgressStateWithSolution(solution string) ProgressState {
	return InProgressState{
		solution: utils.Some(solution),
	}
}
func (s InProgressState) WithSolution(solution string) ProgressState {
	return InProgressState{
		solution: utils.Some(solution),
	}
}

func (s InProgressState) WithNoSolution() ProgressState {
	return InProgressState{
		solution: utils.None[string](),
	}
}

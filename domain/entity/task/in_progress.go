package task

import "github.com/leona-art/task-manager/utils"

type TaskInProgressStatus struct {
	resolution utils.Option[string]
}

func (s TaskInProgressStatus) Status() TaskStatus {
	return InProgress
}

func (s TaskInProgressStatus) Resolution() (value string, ok bool) {
	return s.resolution.Get()
}

func (s TaskInProgressStatus) Candidate() TransitionMap {
	var transitionMap = make(TransitionMap)
	if value, ok := s.resolution.Get(); ok {
		transitionMap[Completed] = func() TaskState {
			return NewTaskCompletedState(value)
		}
	}
	return transitionMap
}

func NewTaskInProgressState() TaskInProgressStatus {
	return TaskInProgressStatus{
		resolution: utils.None[string](),
	}
}

func NewTaskInProgressStateWithResolution(resolution string) TaskInProgressStatus {
	return TaskInProgressStatus{
		resolution: utils.Some(resolution),
	}
}

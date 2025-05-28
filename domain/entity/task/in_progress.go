package task

import "github.com/leona-art/task-manager/utils"

type TaskInProgressStatus struct {
	resolution utils.Option[string]
}

func (s TaskInProgressStatus) Status() string {
	return InProgress
}

func (s TaskInProgressStatus) Resolution() (value string, ok bool) {
	return s.resolution.Get()
}

func (s TaskInProgressStatus) Candidate() map[string]func() TaskStatus {
	value, ok := s.resolution.Get()
	if !ok {
		return map[string]func() TaskStatus{}
	}

	return map[string]func() TaskStatus{
		Completed: func() TaskStatus {
			return NewTaskCompletedStatus(value)
		},
	}
}

func NewTaskInProgressStatus() TaskInProgressStatus {
	return TaskInProgressStatus{
		resolution: utils.None[string](),
	}
}

func NewTaskInProgressStatusWithResolution(resolution string) TaskInProgressStatus {
	return TaskInProgressStatus{
		resolution: utils.Some(resolution),
	}
}

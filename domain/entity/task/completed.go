package task

type TaskCompletedStatus struct {
	resolution string
}

func (s TaskCompletedStatus) Status() TaskStatus {
	return Completed
}

func (s TaskCompletedStatus) Resolution() (value string, ok bool) {
	return s.resolution, true
}

func (s TaskCompletedStatus) Candidate() TransitionMap {
	return TransitionMap{
		InProgress: func() TaskState {
			return NewTaskInProgressStateWithResolution(s.resolution)
		},
	}
}

func NewTaskCompletedState(resolution string) TaskCompletedStatus {
	return TaskCompletedStatus{
		resolution: resolution,
	}
}

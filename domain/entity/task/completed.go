package task

type TaskCompletedStatus struct {
	resolution string
}

func (s TaskCompletedStatus) Status() string {
	return Completed
}

func (s TaskCompletedStatus) Resolution() (value string, ok bool) {
	return s.resolution, true
}

func (s TaskCompletedStatus) Candidate() map[string]func() TaskStatus {
	return map[string]func() TaskStatus{
		InProgress: func() TaskStatus {
			return NewTaskInProgressStatusWithResolution(s.resolution)
		},
	}
}

func NewTaskCompletedStatus(resolution string) TaskCompletedStatus {
	return TaskCompletedStatus{
		resolution: resolution,
	}
}

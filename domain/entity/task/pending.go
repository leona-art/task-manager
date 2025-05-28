package task

type TaskPendingStatus struct{}

func (s TaskPendingStatus) Status() string {
	return Pending
}

func (s TaskPendingStatus) Resolution() (value string, ok bool) {
	return "", false
}
func (s TaskPendingStatus) Candidate() map[string]func() TaskStatus {
	return map[string]func() TaskStatus{
		InProgress: func() TaskStatus {
			return NewTaskInProgressStatus()
		},
	}
}

func NewTaskPendingStatus() TaskPendingStatus {
	return TaskPendingStatus{}
}

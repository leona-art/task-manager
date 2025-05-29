package task

type TaskPendingStatus struct{}

func (s TaskPendingStatus) Status() TaskStatus {
	return Pending
}

func (s TaskPendingStatus) Resolution() (value string, ok bool) {
	return "", false
}
func (s TaskPendingStatus) Candidate() TransitionMap {
	return TransitionMap{
		InProgress: func() TaskState {
			return NewTaskInProgressState()
		},
	}
}

func NewTaskPendingState() TaskPendingStatus {
	return TaskPendingStatus{}
}

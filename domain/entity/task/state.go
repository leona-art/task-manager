package task

type TransitionMap map[TaskStatus]func() TaskState
type TaskState interface {
	Status() TaskStatus
	Resolution() (value string, ok bool)
	Candidate() TransitionMap
}

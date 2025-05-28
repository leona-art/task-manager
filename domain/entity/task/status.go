package task

type TaskStatus interface {
	Status() string
	Resolution() (value string, ok bool)
	Candidate() map[string]func() TaskStatus
}

const (
	Pending    = "pending"
	InProgress = "in progress"
	Completed  = "completed"
)

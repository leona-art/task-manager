package task

type TaskStatus string

const (
	Pending    TaskStatus = "pending"
	InProgress TaskStatus = "in progress"
	Completed  TaskStatus = "completed"
)

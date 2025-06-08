package progress

type ProgressStatus string

const (
	NotStarted ProgressStatus = "not started"
	InProgress ProgressStatus = "in progress"
	Completed  ProgressStatus = "completed"
)

package task

type TaskKind string

const (
	TaskKindProgress TaskKind = "progress"
	TaskKindTodo     TaskKind = "todo"
	TaskKindIssue    TaskKind = "issue"
)

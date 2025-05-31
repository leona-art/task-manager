package task

type Task interface {
	Info() BaseTask
	Kind() TaskKind
}

package task

type Task interface {
	Data() TaskEntity
	Kind() TaskKind
}

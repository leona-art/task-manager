package todo

type TodoState interface {
	Status() TodoStatus
	Switch() TodoState
}

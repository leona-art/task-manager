package todo

type TodoStatus interface {
	Status() string
	Switch() TodoStatus
}

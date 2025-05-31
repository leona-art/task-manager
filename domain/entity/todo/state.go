package todo

type TransitionMap map[TodoStatus]func() TodoState
type TodoState interface {
	Status() TodoStatus
	Candidates() TransitionMap
}

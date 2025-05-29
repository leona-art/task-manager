package todo

type TodoRepository interface {
	Save(todo Todo) error
	GetByID(id TodoId) (value Todo, ok bool, err error)
	Update(todo Todo) error
	Delete(id TodoId) (ok bool, err error)
	List() (todos []Todo, err error)
}

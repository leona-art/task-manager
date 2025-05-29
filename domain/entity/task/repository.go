package task

type taskRepository interface {
	Save(task Task) error
	GetByID(id TaskId) (value Task, ok bool, err error)
	Update(task Task) error
	Delete(id TaskId) (ok bool, err error)
	List() (tasks []Task, err error)
}

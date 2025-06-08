package task

type TaskTitle string

func NewTaskTitle(title string) (TaskTitle, error) {
	if title == "" {
		return TaskTitle(""), nil
	}
	return TaskTitle(title), nil
}

func (t TaskTitle) IsEmpty() bool {
	return t == ""
}
func (t TaskTitle) String() string {
	return string(t)
}
func (t TaskTitle) Equals(other TaskTitle) bool {
	return t == other
}

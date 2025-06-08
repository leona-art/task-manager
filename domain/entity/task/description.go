package task

type TaskDescription string

func NewTaskDescription(description string) (TaskDescription, error) {
	return TaskDescription(description), nil
}

func (t TaskDescription) IsEmpty() bool {
	return t == ""
}
func (t TaskDescription) String() string {
	return string(t)
}
func (t TaskDescription) Equals(other TaskDescription) bool {
	return t == other
}

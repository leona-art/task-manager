package todo

import "time"

type TodoDoneStatus struct {
	at time.Time
}

func (s TodoDoneStatus) Status() string {
	return "done"
}
func (s TodoDoneStatus) Switch() TodoStatus {
	return NewTodoPendingStatus()
}

func (s TodoDoneStatus) At() time.Time {
	return s.at
}

func NewTodoDoneStatus() TodoDoneStatus {
	return TodoDoneStatus{at: time.Now()}
}

package issue

import (
	"fmt"

	"github.com/leona-art/task-manager/domain/entity/task"
)

type IssueTask struct {
	info  task.TaskEntity
	state IssueState
}

func (i *IssueTask) Info() task.TaskEntity {
	return i.info
}
func (i *IssueTask) Kind() task.TaskKind {
	return task.TaskKindIssue
}

func (i *IssueTask) Investigate() error {
	if next, ok := i.state.Candidates()[Investigating]; ok {
		i.state = next()
	} else {
		return fmt.Errorf("cannot start investigating issue task")
	}
	return nil
}

func (i *IssueTask) Resolve() error {
	if next, ok := i.state.Candidates()[Resolving]; ok {
		i.state = next()
	} else {
		return fmt.Errorf("cannot resolve issue task")
	}
	return nil
}

func (i *IssueTask) Close() error {
	if next, ok := i.state.Candidates()[Closed]; ok {
		i.state = next()
	} else {
		return fmt.Errorf("cannot close issue task")
	}
	return nil
}
func (i *IssueTask) SetCause(cause string) error {
	switch state := i.state.(type) {
	case InvestigatingState:
		i.state = state.WithCause(cause)
	default:
		return fmt.Errorf("cannot set cause for issue task in state %T", state)
	}
	return nil
}

func (i *IssueTask) SetSolution(solution string) error {
	switch state := i.state.(type) {
	case ResolvingState:
		i.state = state.WithSolution(solution)
	default:
		return fmt.Errorf("cannot set solution for issue task in state %T", state)
	}
	return nil
}

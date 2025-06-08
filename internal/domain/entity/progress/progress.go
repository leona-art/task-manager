package progress

import (
	"fmt"

	"github.com/leona-art/task-manager/internal/domain/entity/task"
)

type ProgressTask struct {
	info  task.TaskEntity
	state ProgressState
}

func NewProgressTask(title, description string) (*ProgressTask, error) {
	data, err := task.NewTaskEntity(title, description)
	if err != nil {
		return nil, fmt.Errorf("failed to create task entity: %w", err)
	}
	return &ProgressTask{
		info:  data,
		state: NewNotStartedState(),
	}, nil
}

func (p *ProgressTask) Info() task.TaskEntity {
	return p.info
}

func (p *ProgressTask) Kind() task.TaskKind {
	return task.TaskKindProgress
}

func (p *ProgressTask) Start() error {
	if next, ok := p.state.Candidates()[InProgress]; ok {
		p.state = next()
	} else {
		return fmt.Errorf("cannot start progress task")
	}
	return nil
}

func (p *ProgressTask) SetSolution(solution string) error {
	switch state := p.state.(type) {
	case InProgressState:
		p.state = state.WithSolution(solution)
	default:
		return fmt.Errorf("cannot set solution for progress task in state %T", state)
	}
	return nil
}

func (p *ProgressTask) Complete() error {
	if next, ok := p.state.Candidates()[Completed]; ok {
		p.state = next()
	} else {
		return fmt.Errorf("cannot complete progress task")
	}
	return nil
}

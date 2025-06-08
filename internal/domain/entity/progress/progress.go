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

func NewProgressTaskFromDto(dto ProgressTaskDto) (*ProgressTask, error) {
	id, err := task.NewTaskIdFromString(dto.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID: %v", err)
	}
	data := task.TaskEntity{
		ID:          id,
		Title:       dto.Title,
		Description: dto.Description,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
	if data.IsValid() == false {
		return nil, fmt.Errorf("invalid task entity data")
	}

	var state ProgressState
	switch dto.Status {
	case string(NotStarted):
		state = NewNotStartedState()
	case string(InProgress):
		if solution, ok := dto.Solution.Get(); ok {
			state = NewInProgressStateWithSolution(solution)
		} else {
			state = NewInProgressState()
		}
	case string(Completed):
		if solution, ok := dto.Solution.Get(); ok {
			state = NewCompletedState(solution)
		} else {
			return nil, fmt.Errorf("solution is required for completed state")
		}
	default:
		return nil, fmt.Errorf("invalid task state: %s", dto.Status)
	}

	return &ProgressTask{
		info:  data,
		state: state,
	}, nil
}

func (p *ProgressTask) Data() task.TaskEntity {
	return p.info
}

func (p *ProgressTask) Kind() task.TaskKind {
	return task.TaskKindProgress
}

func (p *ProgressTask) State() ProgressState {
	return p.state
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

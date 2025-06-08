package usecase

import (
	"context"

	"github.com/leona-art/task-manager/internal/adaptor/gateway"
	"github.com/leona-art/task-manager/internal/domain/entity/progress"
	"github.com/leona-art/task-manager/internal/domain/entity/task"
)

type ProgressUseCase struct {
	repository gateway.ProgressRepository
}

func NewProgressUseCase(repository gateway.ProgressRepository) *ProgressUseCase {
	return &ProgressUseCase{
		repository: repository,
	}
}

func (uc *ProgressUseCase) CreateProgress(ctx context.Context, title, description string) (*progress.ProgressTask, error) {
	progressTask, err := progress.NewProgressTask(title, description)
	if err != nil {
		return nil, err
	}
	if err := uc.repository.Create(ctx, *progressTask); err != nil {
		return nil, err
	}
	return progressTask, nil
}

func (uc *ProgressUseCase) GetProgressByID(ctx context.Context, id string) (*progress.ProgressTask, bool, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return nil, false, err
	}
	progressTask, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return nil, false, err
	}
	return &progressTask, ok, nil
}

func (uc *ProgressUseCase) ListProgress(ctx context.Context) ([]progress.ProgressTask, error) {
	progressTasks, err := uc.repository.List(ctx)
	if err != nil {
		return nil, err
	}
	return progressTasks, nil
}

func (uc *ProgressUseCase) StartProgress(ctx context.Context, id string) (*progress.ProgressTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return nil, err
	}
	progressTask, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil // Task not found
	}
	progressTask.Start()
	if err := uc.repository.Save(ctx, progressTask); err != nil {
		return nil, err
	}
	return &progressTask, nil
}

func (uc *ProgressUseCase) CompleteProgress(ctx context.Context, id string) (*progress.ProgressTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return nil, err
	}
	progressTask, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil // Task not found
	}
	progressTask.Complete()
	if err := uc.repository.Save(ctx, progressTask); err != nil {
		return nil, err
	}
	return &progressTask, nil
}

func (uc *ProgressUseCase) DeleteProgress(ctx context.Context, id string) (bool, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return false, err
	}
	ok, err := uc.repository.Delete(ctx, taskId)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (uc *ProgressUseCase) SetSolution(ctx context.Context, id string, solution string) (*progress.ProgressTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return nil, err
	}
	progressTask, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil // Task not found
	}
	progressTask.SetSolution(solution)
	if err := uc.repository.Save(ctx, progressTask); err != nil {
		return nil, err
	}
	return &progressTask, nil
}

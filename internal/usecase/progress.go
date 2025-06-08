package usecase

import (
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

func (uc *ProgressUseCase) CreateProgress(title, description string) error {
	progressTask, err := progress.NewProgressTask(title, description)
	if err != nil {
		return err
	}
	if err := uc.repository.Create(*progressTask); err != nil {
		return err
	}
	return nil
}

func (uc *ProgressUseCase) GetProgressByID(id string) (progress.ProgressTask, bool, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return progress.ProgressTask{}, false, err
	}
	progressTask, ok, err := uc.repository.Get(taskId)
	if err != nil {
		return progress.ProgressTask{}, false, err
	}
	return progressTask, ok, nil
}

func (uc *ProgressUseCase) ListProgress() ([]progress.ProgressTask, error) {
	progressTasks, err := uc.repository.List()
	if err != nil {
		return nil, err
	}
	return progressTasks, nil
}

func (uc *ProgressUseCase) StartProgress(id string) (progress.ProgressTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return progress.ProgressTask{}, err
	}
	progressTask, ok, err := uc.repository.Get(taskId)
	if err != nil {
		return progress.ProgressTask{}, err
	}
	if !ok {
		return progress.ProgressTask{}, nil // Task not found
	}
	progressTask.Start()
	if err := uc.repository.Save(progressTask); err != nil {
		return progress.ProgressTask{}, err
	}
	return progressTask, nil
}

func (uc *ProgressUseCase) CompleteProgress(id string) (progress.ProgressTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return progress.ProgressTask{}, err
	}
	progressTask, ok, err := uc.repository.Get(taskId)
	if err != nil {
		return progress.ProgressTask{}, err
	}
	if !ok {
		return progress.ProgressTask{}, nil // Task not found
	}
	progressTask.Complete()
	if err := uc.repository.Save(progressTask); err != nil {
		return progress.ProgressTask{}, err
	}
	return progressTask, nil
}

func (uc *ProgressUseCase) DeleteProgress(id string) (bool, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return false, err
	}
	ok, err := uc.repository.Delete(taskId)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (uc *ProgressUseCase) SetSolution(id string, solution string) (progress.ProgressTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return progress.ProgressTask{}, err
	}
	progressTask, ok, err := uc.repository.Get(taskId)
	if err != nil {
		return progress.ProgressTask{}, err
	}
	if !ok {
		return progress.ProgressTask{}, nil // Task not found
	}
	progressTask.SetSolution(solution)
	if err := uc.repository.Save(progressTask); err != nil {
		return progress.ProgressTask{}, err
	}
	return progressTask, nil
}

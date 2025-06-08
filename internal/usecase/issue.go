package usecase

import (
	"context"
	"fmt"

	"github.com/leona-art/task-manager/internal/adaptor/gateway"
	"github.com/leona-art/task-manager/internal/domain/entity/issue"
	"github.com/leona-art/task-manager/internal/domain/entity/task"
)

type IssueUseCase struct {
	repository gateway.IssueRepository
}

func NewIssueUseCase(repository gateway.IssueRepository) *IssueUseCase {
	return &IssueUseCase{
		repository: repository,
	}
}

func (uc *IssueUseCase) CreateIssue(ctx context.Context, title, description string) (*issue.IssueTask, error) {
	issueTask, err := issue.NewIssueTask(title, description)
	if err != nil {
		return nil, err
	}
	if err := uc.repository.Create(ctx, *issueTask); err != nil {
		return nil, err
	}
	return issueTask, nil
}

func (uc *IssueUseCase) GetIssueByID(ctx context.Context, id string) (*issue.IssueTask, bool, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return nil, false, err
	}
	issueTask, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return nil, false, err
	}
	return &issueTask, ok, nil
}

func (uc *IssueUseCase) ListIssues(ctx context.Context) ([]issue.IssueTask, error) {
	issueTasks, err := uc.repository.List(ctx)
	if err != nil {
		return nil, err
	}
	return issueTasks, nil
}

func (uc *IssueUseCase) InvestigateIssue(ctx context.Context, id string) (*issue.IssueTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return nil, err
	}
	issueTask, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("issue task with id %s not found", id)
	}

	if err := issueTask.Investigate(); err != nil {
		return nil, err
	}
	if err := uc.repository.Save(ctx, issueTask); err != nil {
		return nil, err
	}
	return &issueTask, nil
}

func (uc *IssueUseCase) ResolveIssue(ctx context.Context, id string) (*issue.IssueTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return nil, err
	}
	issueTask, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("issue task with id %s not found", id)
	}

	if err := issueTask.Resolve(); err != nil {
		return nil, err
	}
	if err := uc.repository.Save(ctx, issueTask); err != nil {
		return nil, err
	}
	return &issueTask, nil
}

func (uc *IssueUseCase) CloseIssue(ctx context.Context, id string) (issue.IssueTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return issue.IssueTask{}, err
	}
	issueTask, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return issue.IssueTask{}, err
	}
	if !ok {
		return issue.IssueTask{}, fmt.Errorf("issue task with id %s not found", id)
	}

	if err := issueTask.Close(); err != nil {
		return issue.IssueTask{}, err
	}
	if err := uc.repository.Save(ctx, issueTask); err != nil {
		return issue.IssueTask{}, err
	}
	return issueTask, nil
}

func (uc *IssueUseCase) SetCause(ctx context.Context, id string, cause string) (issue.IssueTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return issue.IssueTask{}, err
	}
	issueTask, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return issue.IssueTask{}, err
	}
	if !ok {
		return issue.IssueTask{}, fmt.Errorf("issue task with id %s not found", id)
	}

	if err := issueTask.SetCause(cause); err != nil {
		return issue.IssueTask{}, err
	}
	if err := uc.repository.Save(ctx, issueTask); err != nil {
		return issue.IssueTask{}, err
	}
	return issueTask, nil
}

func (uc *IssueUseCase) SetSolution(ctx context.Context, id string, solution string) (issue.IssueTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return issue.IssueTask{}, err
	}
	issueTask, ok, err := uc.repository.Get(ctx, taskId)
	if err != nil {
		return issue.IssueTask{}, err
	}
	if !ok {
		return issue.IssueTask{}, fmt.Errorf("issue task with id %s not found", id)
	}

	if err := issueTask.SetSolution(solution); err != nil {
		return issue.IssueTask{}, err
	}
	if err := uc.repository.Save(ctx, issueTask); err != nil {
		return issue.IssueTask{}, err
	}
	return issueTask, nil
}

func (uc *IssueUseCase) DeleteIssue(ctx context.Context, id string) (bool, error) {
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

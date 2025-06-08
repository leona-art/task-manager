package usecase

import (
	"fmt"

	"github.com/leona-art/task-manager/domain/entity/issue"
	"github.com/leona-art/task-manager/domain/entity/task"
)

type IssueUseCase struct {
	repository issue.IssueRepository
}

func NewIssueUseCase(repository issue.IssueRepository) *IssueUseCase {
	return &IssueUseCase{
		repository: repository,
	}
}

func (uc *IssueUseCase) CreateIssue(title, description string) (*issue.IssueTask, error) {
	issueTask, err := issue.NewIssueTask(title, description)
	if err != nil {
		return nil, err
	}
	if err := uc.repository.Create(*issueTask); err != nil {
		return nil, err
	}
	return issueTask, nil
}

func (uc *IssueUseCase) GetIssueByID(id string) (*issue.IssueTask, bool, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return nil, false, err
	}
	issueTask, ok, err := uc.repository.Get(taskId)
	if err != nil {
		return nil, false, err
	}
	return &issueTask, ok, nil
}

func (uc *IssueUseCase) ListIssues() ([]issue.IssueTask, error) {
	issueTasks, err := uc.repository.List()
	if err != nil {
		return nil, err
	}
	return issueTasks, nil
}

func (uc *IssueUseCase) InvestigateIssue(id string) (*issue.IssueTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return nil, err
	}
	issueTask, ok, err := uc.repository.Get(taskId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("issue task with id %s not found", id)
	}

	if err := issueTask.Investigate(); err != nil {
		return nil, err
	}
	if err := uc.repository.Save(issueTask); err != nil {
		return nil, err
	}
	return &issueTask, nil
}

func (uc *IssueUseCase) ResolveIssue(id string) (*issue.IssueTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return nil, err
	}
	issueTask, ok, err := uc.repository.Get(taskId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("issue task with id %s not found", id)
	}

	if err := issueTask.Resolve(); err != nil {
		return nil, err
	}
	if err := uc.repository.Save(issueTask); err != nil {
		return nil, err
	}
	return &issueTask, nil
}

func (uc *IssueUseCase) CloseIssue(id string) (issue.IssueTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return issue.IssueTask{}, err
	}
	issueTask, ok, err := uc.repository.Get(taskId)
	if err != nil {
		return issue.IssueTask{}, err
	}
	if !ok {
		return issue.IssueTask{}, fmt.Errorf("issue task with id %s not found", id)
	}

	if err := issueTask.Close(); err != nil {
		return issue.IssueTask{}, err
	}
	if err := uc.repository.Save(issueTask); err != nil {
		return issue.IssueTask{}, err
	}
	return issueTask, nil
}

func (uc *IssueUseCase) SetCause(id string, cause string) (issue.IssueTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return issue.IssueTask{}, err
	}
	issueTask, ok, err := uc.repository.Get(taskId)
	if err != nil {
		return issue.IssueTask{}, err
	}
	if !ok {
		return issue.IssueTask{}, fmt.Errorf("issue task with id %s not found", id)
	}

	if err := issueTask.SetCause(cause); err != nil {
		return issue.IssueTask{}, err
	}
	if err := uc.repository.Save(issueTask); err != nil {
		return issue.IssueTask{}, err
	}
	return issueTask, nil
}

func (uc *IssueUseCase) SetSolution(id string, solution string) (issue.IssueTask, error) {
	taskId, err := task.NewTaskIdFromString(id)
	if err != nil {
		return issue.IssueTask{}, err
	}
	issueTask, ok, err := uc.repository.Get(taskId)
	if err != nil {
		return issue.IssueTask{}, err
	}
	if !ok {
		return issue.IssueTask{}, fmt.Errorf("issue task with id %s not found", id)
	}

	if err := issueTask.SetSolution(solution); err != nil {
		return issue.IssueTask{}, err
	}
	if err := uc.repository.Save(issueTask); err != nil {
		return issue.IssueTask{}, err
	}
	return issueTask, nil
}

func (uc *IssueUseCase) DeleteIssue(id string) (bool, error) {
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

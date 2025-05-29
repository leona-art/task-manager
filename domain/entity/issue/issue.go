package issue

import (
	"fmt"

	"github.com/leona-art/task-manager/domain/entity/task_info"
)

type Issue struct {
	Info   task_info.TaskInfo
	Status IssueState
}

func NewIssue(title, description string) (Issue, error) {
	info, err := task_info.NewTaskInfo(title, description)
	if err != nil {
		return Issue{}, fmt.Errorf("failed to generate UUID: %w", err)
	}
	return Issue{
		Info:   info,
		Status: NewIssuePendingState(),
	}, nil
}
func (i *Issue) StartResearching() error {
	candidate := i.Status.Candidate()
	if progress, ok := candidate[Researching]; ok {
		i.Status = progress()
	} else {
		return fmt.Errorf("cannot open issue for research: no candidate status found")
	}
	return nil
}

func (i *Issue) StartResolution() error {
	candidate := i.Status.Candidate()
	if resolve, ok := candidate[Resolving]; ok {
		i.Status = resolve()
	} else {
		return fmt.Errorf("cannot open issue for resolution: no candidate status found")
	}
	return nil
}

func (i *Issue) Resolve(resolution string) error {
	candidate := i.Status.Candidate()
	if resolved, ok := candidate[Resolved]; ok {
		i.Status = resolved()
	} else {
		return fmt.Errorf("cannot resolve issue: no candidate status found")
	}
	return nil
}

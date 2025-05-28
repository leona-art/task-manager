package issue

import "fmt"

type Issue struct {
	ID          string
	Title       string
	Description string
	Status      IssueStatus
}

func NewIssue(id, title, description string) Issue {
	return Issue{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      NewIssuePendingStatus(),
	}
}
func (i *Issue) OpenResearch() error {
	candidate := i.Status.Candidate()
	if progress, ok := candidate[Researching]; ok {
		i.Status = progress()
	} else {
		return fmt.Errorf("cannot open issue for research: no candidate status found")
	}
	return nil
}

func (i *Issue) OpenResolve() error {
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

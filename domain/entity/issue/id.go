package issue

import "github.com/google/uuid"

type IssueId string

func NewIssueId() (IssueId, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return IssueId(id.String()), nil
}

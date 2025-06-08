package issue

import "github.com/leona-art/task-manager/internal/utils"

type IssueTaskDto struct {
	ID          string
	Title       string
	Description string
	State       string
	Cause       utils.Option[string]
	Solution    utils.Option[string]
	CreatedAt   string
	UpdatedAt   string
}

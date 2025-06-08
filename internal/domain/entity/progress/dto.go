package progress

import (
	"time"

	"github.com/leona-art/task-manager/internal/utils"
)

type ProgressTaskDto struct {
	ID          string
	Title       string
	Description string
	State       string
	Solution    utils.Option[string]
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

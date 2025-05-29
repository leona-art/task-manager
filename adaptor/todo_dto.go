package adaptor

import (
	"github.com/leona-art/task-manager/domain/entity/todo"
	workspacev1 "github.com/leona-art/task-manager/gen/workspace/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToTodoDTO(t todo.Todo) (dto *workspacev1.Todo) {
	dto = &workspacev1.Todo{}
	dto.Id = t.ID.String()
	dto.Title = t.Info.Title
	dto.Description = t.Info.Description
	dto.CreatedAt = timestamppb.New(t.Info.CreatedAt)
	dto.UpdatedAt = timestamppb.New(t.Info.UpdatedAt)

	switch t.State.Status() {
	case todo.Done:
		dto.Status = workspacev1.TodoStatus_TODO_STATUS_DONE
	case todo.Pending:
		dto.Status = workspacev1.TodoStatus_TODO_STATUS_PENDING
	}
	return
}

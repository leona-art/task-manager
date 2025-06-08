package controller

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/leona-art/task-manager/domain/entity/todo"
	workspacev1 "github.com/leona-art/task-manager/gen/workspace/v1"
	"github.com/leona-art/task-manager/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WorkSpaceController struct {
	TodoUsecase *usecase.TodoUseCase
}

// CreateTodo implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) CreateTodo(ctx context.Context, req *connect.Request[workspacev1.CreateTodoRequest]) (*connect.Response[workspacev1.CreateTodoResponse], error) {
	slog.Info("CreateTodo called", "title", req.Msg.Title, "description", req.Msg.Description)
	t, err := w.TodoUsecase.CreateTodo(ctx, req.Msg.Title, req.Msg.Description)
	if err != nil {
		slog.Error("failed to create todo", "error", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	pbTodo := toPbTodo(&t)
	return connect.NewResponse(&workspacev1.CreateTodoResponse{
		Todo: pbTodo,
	}), nil
}

// DoTodo implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) DoTodo(ctx context.Context, req *connect.Request[workspacev1.DoTodoRequest]) (*connect.Response[workspacev1.DoTodoResponse], error) {
	slog.Info("DoTodo called", "id", req.Msg.Id)
	t, err := w.TodoUsecase.MarkAsDoneTodo(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to do todo: %w", err))
	}
	pbTodo := toPbTodo(&t)
	return connect.NewResponse(&workspacev1.DoTodoResponse{
		Todo: pbTodo,
	}), nil
}

// GetTodo implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) GetTodo(ctx context.Context, req *connect.Request[workspacev1.GetTodoRequest]) (*connect.Response[workspacev1.GetTodoResponse], error) {
	t, ok, err := w.TodoUsecase.GetTodoByID(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get todo by id: %w", err))
	}
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("todo with id %s not found", req.Msg.Id))
	}
	pbTodo := toPbTodo(&t)
	return connect.NewResponse(&workspacev1.GetTodoResponse{
		Todo: pbTodo,
	}), nil
}

// ListTodos implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) ListTodos(ctx context.Context, req *connect.Request[workspacev1.ListTodosRequest]) (*connect.Response[workspacev1.ListTodosResponse], error) {
	todos, err := w.TodoUsecase.ListTodos(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list todos: %w", err))
	}
	pbTodos := make([]*workspacev1.Todo, len(todos))
	for i, t := range todos {
		pbTodos[i] = toPbTodo(&t)
	}
	return connect.NewResponse(&workspacev1.ListTodosResponse{
		Todos: pbTodos,
	}), nil
}

// UndoneTodo implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) UndoneTodo(ctx context.Context, req *connect.Request[workspacev1.UndoneTodoRequest]) (*connect.Response[workspacev1.UndoneTodoResponse], error) {
	t, err := w.TodoUsecase.RevertTodo(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to undone todo: %w", err))
	}
	pbTodo := toPbTodo(&t)
	return connect.NewResponse(&workspacev1.UndoneTodoResponse{
		Todo: pbTodo,
	}), nil
}

func toDomainTodo(todoPb *workspacev1.Todo) (*todo.TodoTask, error) {
	if todoPb == nil {
		return nil, fmt.Errorf("todo cannot be nil")
	}
	todoDTO := todo.TodoTaskDto{
		ID:          todoPb.Id,
		Title:       todoPb.Title,
		Description: todoPb.Description,
		CreatedAt:   todoPb.CreatedAt.AsTime(),
		UpdatedAt:   todoPb.UpdatedAt.AsTime(),
		Status:      todoPb.Status.String(),
	}
	t, err := todo.NewTodoTaskFromDto(todoDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to convert todo to domain: %w", err)
	}
	return &t, nil
}

func toPbTodo(t *todo.TodoTask) (pb *workspacev1.Todo) {
	pb = &workspacev1.Todo{}
	pb.Id = t.Data().ID.String()
	pb.Title = t.Data().Title
	pb.Description = t.Data().Description
	pb.CreatedAt = timestamppb.New(t.Data().CreatedAt)
	pb.UpdatedAt = timestamppb.New(t.Data().UpdatedAt)

	switch t.State().Status() {
	case todo.Done:
		pb.Status = workspacev1.TodoStatus_TODO_STATUS_DONE
	case todo.Pending:
		pb.Status = workspacev1.TodoStatus_TODO_STATUS_PENDING
	}
	return
}

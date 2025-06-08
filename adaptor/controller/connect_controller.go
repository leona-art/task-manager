package controller

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/leona-art/task-manager/adaptor"
	workspacev1 "github.com/leona-art/task-manager/gen/workspace/v1"
	"github.com/leona-art/task-manager/usecase"
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
	dto := adaptor.ToTodoDTO(t)
	return connect.NewResponse(&workspacev1.CreateTodoResponse{
		Todo: dto,
	}), nil
}

// DoTodo implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) DoTodo(ctx context.Context, req *connect.Request[workspacev1.DoTodoRequest]) (*connect.Response[workspacev1.DoTodoResponse], error) {
	t, err := w.TodoUsecase.MarkAsDoneTodo(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to do todo: %w", err))
	}
	dto := adaptor.ToTodoDTO(t)
	return connect.NewResponse(&workspacev1.DoTodoResponse{
		Todo: dto,
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
	dto := adaptor.ToTodoDTO(t)
	return connect.NewResponse(&workspacev1.GetTodoResponse{
		Todo: dto,
	}), nil
}

// ListTodos implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) ListTodos(ctx context.Context, req *connect.Request[workspacev1.ListTodosRequest]) (*connect.Response[workspacev1.ListTodosResponse], error) {
	todos, err := w.TodoUsecase.ListTodos(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list todos: %w", err))
	}
	dtos := make([]*workspacev1.Todo, len(todos))
	for i, t := range todos {
		dtos[i] = adaptor.ToTodoDTO(t)
	}
	return connect.NewResponse(&workspacev1.ListTodosResponse{
		Todos: dtos,
	}), nil
}

// UndoneTodo implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) UndoneTodo(ctx context.Context, req *connect.Request[workspacev1.UndoneTodoRequest]) (*connect.Response[workspacev1.UndoneTodoResponse], error) {
	t, err := w.TodoUsecase.RevertTodo(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to undone todo: %w", err))
	}
	dto := adaptor.ToTodoDTO(t)
	return connect.NewResponse(&workspacev1.UndoneTodoResponse{
		Todo: dto,
	}), nil
}

package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	workspacev1 "github.com/leona-art/task-manager/gen/workspace/v1"
	"github.com/leona-art/task-manager/gen/workspace/v1/workspacev1connect"
	"github.com/leona-art/task-manager/infra/inmemory"
	"github.com/leona-art/task-manager/usecase"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type WorkSpaceServer struct {
	todoUsecase *usecase.TodoUseCase
}

// CreateTodo implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceServer) CreateTodo(ctx context.Context, req *connect.Request[workspacev1.CreateTodoRequest]) (*connect.Response[workspacev1.CreateTodoResponse], error) {
	// t, err := w.todoUsecase.CreateTodo(req.Msg.Title, req.Msg.Description)
	// if err != nil {
	// 	return nil, err
	// }
	// dto := adaptor.ToTodoDTO(t)
	// return connect.NewResponse(&workspacev1.CreateTodoResponse{
	// 	Todo: dto,
	// }), nil
	panic("CreateTodo not implemented yet")
}

// DoTodo implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceServer) DoTodo(ctx context.Context, req *connect.Request[workspacev1.DoTodoRequest]) (*connect.Response[workspacev1.DoTodoResponse], error) {
	// todoId, err := todo.NewTodoIdFromString(req.Msg.Id)
	// if err != nil {
	// 	return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid todo id: %w", err))
	// }
	// t, err := w.todoUsecase.DoTodo(todoId)
	// if err != nil {
	// 	return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to do todo: %w", err))
	// }
	// dto := adaptor.ToTodoDTO(t)
	// return connect.NewResponse(&workspacev1.DoTodoResponse{
	// 	Todo: dto,
	// }), nil
	panic("DoTodo not implemented yet")
}

// GetTodo implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceServer) GetTodo(ctx context.Context, req *connect.Request[workspacev1.GetTodoRequest]) (*connect.Response[workspacev1.GetTodoResponse], error) {
	// todoId, err := todo.NewTodoIdFromString(req.Msg.Id)
	// if err != nil {
	// 	return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid todo id: %w", err))
	// }
	// t, ok, err := w.todoUsecase.GetTodoByID(todoId)
	// if err != nil {
	// 	return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get todo by id: %w", err))
	// }
	// if !ok {
	// 	return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("todo with id %s not found", todoId.String()))
	// }
	// dto := adaptor.ToTodoDTO(t)
	// return connect.NewResponse(&workspacev1.GetTodoResponse{
	// 	Todo: dto,
	// }), nil
	panic("GetTodo not implemented yet")
}

// ListTodos implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceServer) ListTodos(context.Context, *connect.Request[workspacev1.ListTodosRequest]) (*connect.Response[workspacev1.ListTodosResponse], error) {
	// todos, err := w.todoUsecase.ListTodos()
	// if err != nil {
	// 	return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list todos: %w", err))
	// }
	// dtos := make([]*workspacev1.Todo, len(todos))
	// for i, t := range todos {
	// 	dtos[i] = adaptor.ToTodoDTO(t)
	// }
	// return connect.NewResponse(&workspacev1.ListTodosResponse{
	// 	Todos: dtos,
	// }), nil
	panic("ListTodos not implemented yet")
}

// UndoneTodo implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceServer) UndoneTodo(ctx context.Context, req *connect.Request[workspacev1.UndoneTodoRequest]) (*connect.Response[workspacev1.UndoneTodoResponse], error) {
	// todoId, err := todo.NewTodoIdFromString(req.Msg.Id)
	// if err != nil {
	// 	return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid todo id: %w", err))
	// }
	// t, err := w.todoUsecase.UndoTodo(todoId)
	// if err != nil {
	// 	return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to undone todo: %w", err))
	// }
	// dto := adaptor.ToTodoDTO(t)
	// return connect.NewResponse(&workspacev1.UndoneTodoResponse{
	// 	Todo: dto,
	// }), nil
	panic("UndoneTodo not implemented yet")
}

func main() {
	todoRepo := inmemory.NewInMemoryTodoRepository()
	todoUsecase := usecase.NewTodoUseCase(todoRepo)
	workspace := &WorkSpaceServer{
		todoUsecase: todoUsecase,
	}
	mux := http.NewServeMux()
	path, handler := workspacev1connect.NewWorkspaceServiceHandler(workspace)
	mux.Handle(path, handler)
	http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

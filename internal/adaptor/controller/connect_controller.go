package controller

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	workspacev1 "github.com/leona-art/task-manager/gen/workspace/v1"
	"github.com/leona-art/task-manager/internal/domain/entity/progress"
	"github.com/leona-art/task-manager/internal/domain/entity/todo"
	"github.com/leona-art/task-manager/internal/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WorkSpaceController struct {
	TodoUsecase     *usecase.TodoUseCase
	ProgressUsecase *usecase.ProgressUseCase
}

// CompleteProgress implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) CompleteProgress(ctx context.Context, req *connect.Request[workspacev1.CompleteProgressRequest]) (*connect.Response[workspacev1.CompleteProgressResponse], error) {
	progressTask, err := w.ProgressUsecase.CompleteProgress(req.Msg.Id)
	if err != nil {
		slog.Error("failed to complete progress", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to complete progress: %w", err))
	}
	pbProgress, err := toPbProgress(progressTask)
	if err != nil {
		slog.Error("failed to convert progress to protobuf", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert progress to protobuf: %w", err))
	}
	return connect.NewResponse(&workspacev1.CompleteProgressResponse{
		Progress: pbProgress,
	}), nil
}

// CreateProgress implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) CreateProgress(ctx context.Context, req *connect.Request[workspacev1.CreateProgressRequest]) (*connect.Response[workspacev1.CreateProgressResponse], error) {
	progressTask, err := w.ProgressUsecase.CreateProgress(req.Msg.Title, req.Msg.Description)
	if err != nil {
		slog.Error("failed to create progress", "error", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	pbProgress, err := toPbProgress(progressTask)
	if err != nil {
		slog.Error("failed to convert progress to protobuf", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert progress to protobuf: %w", err))
	}
	return connect.NewResponse(&workspacev1.CreateProgressResponse{
		Progress: pbProgress,
	}), nil
}

// DeleteProgress implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) DeleteProgress(ctx context.Context, req *connect.Request[workspacev1.DeleteProgressRequest]) (*connect.Response[workspacev1.DeleteProgressResponse], error) {
	ok, err := w.ProgressUsecase.DeleteProgress(req.Msg.Id)
	if err != nil {
		slog.Error("failed to delete progress", "error", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("progress with id %s not found", req.Msg.Id))
	}
	if err != nil {
		slog.Error("failed to convert progress to protobuf", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert progress to protobuf: %w", err))
	}
	return connect.NewResponse(&workspacev1.DeleteProgressResponse{}), nil
}

// DeleteTodo implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) DeleteTodo(ctx context.Context, req *connect.Request[workspacev1.DeleteTodoRequest]) (*connect.Response[workspacev1.DeleteTodoResponse], error) {
	ok, err := w.TodoUsecase.DeleteTodo(ctx, req.Msg.Id)
	if err != nil {
		slog.Error("failed to delete todo", "error", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("todo with id %s not found", req.Msg.Id))
	}
	return connect.NewResponse(&workspacev1.DeleteTodoResponse{}), nil
}

// GetProgress implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) GetProgress(ctx context.Context, req *connect.Request[workspacev1.GetProgressRequest]) (*connect.Response[workspacev1.GetProgressResponse], error) {
	progressTask, ok, err := w.ProgressUsecase.GetProgressByID(req.Msg.Id)
	if err != nil {
		slog.Error("failed to get progress by id", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get progress by id: %w", err))
	}
	if !ok {
		slog.Error("progress not found", "id", req.Msg.Id)
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("progress with id %s not found", req.Msg.Id))
	}
	pbProgress, err := toPbProgress(progressTask)
	if err != nil {
		slog.Error("failed to convert progress to protobuf", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert progress to protobuf: %w", err))
	}
	return connect.NewResponse(&workspacev1.GetProgressResponse{
		Progress: pbProgress,
	}), nil
}

// ListProgresses implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) ListProgresses(ctx context.Context, req *connect.Request[workspacev1.ListProgressesRequest]) (*connect.Response[workspacev1.ListProgressesResponse], error) {
	progressTasks, err := w.ProgressUsecase.ListProgress()
	if err != nil {
		slog.Error("failed to list progresses", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list progresses: %w", err))
	}
	pbProgresses := make([]*workspacev1.Progress, len(progressTasks))
	for i, pt := range progressTasks {
		pbProgress, err := toPbProgress(&pt)
		if err != nil {
			slog.Error("failed to convert progress to protobuf", "error", err)
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert progress to protobuf: %w", err))
		}
		pbProgresses[i] = pbProgress
	}
	return connect.NewResponse(&workspacev1.ListProgressesResponse{
		Progresses: pbProgresses,
	}), nil
}

// SetProgressSolution implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) SetProgressSolution(ctx context.Context, req *connect.Request[workspacev1.SetProgressSolutionRequest]) (*connect.Response[workspacev1.SetProgressSolutionResponse], error) {
	progressTask, err := w.ProgressUsecase.SetSolution(req.Msg.Id, req.Msg.Solution)
	if err != nil {
		slog.Error("failed to set progress solution", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to set progress solution: %w", err))
	}
	pbProgress, err := toPbProgress(progressTask)
	if err != nil {
		slog.Error("failed to convert progress to protobuf", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert progress to protobuf: %w", err))
	}
	return connect.NewResponse(&workspacev1.SetProgressSolutionResponse{
		Progress: pbProgress,
	}), nil
}

// StartProgress implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) StartProgress(ctx context.Context, req *connect.Request[workspacev1.StartProgressRequest]) (*connect.Response[workspacev1.StartProgressResponse], error) {
	progressTask, err := w.ProgressUsecase.StartProgress(req.Msg.Id)
	if err != nil {
		slog.Error("failed to start progress", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to start progress: %w", err))
	}
	pbProgress, err := toPbProgress(progressTask)
	if err != nil {
		slog.Error("failed to convert progress to protobuf", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert progress to protobuf: %w", err))
	}
	return connect.NewResponse(&workspacev1.StartProgressResponse{
		Progress: pbProgress,
	}), nil
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

func toPbProgress(progressTask *progress.ProgressTask) (pb *workspacev1.Progress, err error) {
	pb = &workspacev1.Progress{}
	pb.Id = progressTask.Data().ID.String()
	pb.Title = progressTask.Data().Title
	pb.Description = progressTask.Data().Description
	pb.CreatedAt = timestamppb.New(progressTask.Data().CreatedAt)
	pb.UpdatedAt = timestamppb.New(progressTask.Data().UpdatedAt)

	switch s := progressTask.State().(type) {
	case progress.InProgressState:
		pb.State = &workspacev1.ProgressState{
			State: &workspacev1.ProgressState_NotStarted{},
		}
	case progress.CompletedState:
		s.Solution()
	}

	switch progressTask.State().Status() {
	case progress.NotStarted:
		pb.State = &workspacev1.ProgressState{
			State: &workspacev1.ProgressState_NotStarted{},
		}
	case progress.InProgress:
		if solution, ok := progressTask.State().Solution(); ok {
			pb.State = &workspacev1.ProgressState{
				State: &workspacev1.ProgressState_InProgress{
					InProgress: &workspacev1.ProgressInProgressState{
						Solution: &solution,
					},
				},
			}
		} else {
			pb.State = &workspacev1.ProgressState{
				State: &workspacev1.ProgressState_InProgress{
					InProgress: &workspacev1.ProgressInProgressState{
						Solution: nil,
					},
				},
			}
		}
	case progress.Completed:
		if solution, ok := progressTask.State().Solution(); ok {
			pb.State = &workspacev1.ProgressState{
				State: &workspacev1.ProgressState_Completed{
					Completed: &workspacev1.ProgressCompletedState{
						Solution: solution,
					},
				},
			}
		} else {
			return nil, fmt.Errorf("solution is required for completed state")
		}
	default:
		return nil, fmt.Errorf("unknown progress state: %T", progressTask.State())
	}
	return
}

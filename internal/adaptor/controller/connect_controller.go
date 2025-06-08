package controller

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	workspacev1 "github.com/leona-art/task-manager/gen/workspace/v1"
	"github.com/leona-art/task-manager/internal/domain/entity/issue"
	"github.com/leona-art/task-manager/internal/domain/entity/progress"
	"github.com/leona-art/task-manager/internal/domain/entity/todo"
	"github.com/leona-art/task-manager/internal/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WorkSpaceController struct {
	TodoUsecase     *usecase.TodoUseCase
	ProgressUsecase *usecase.ProgressUseCase
	IssueUsecase    *usecase.IssueUseCase
}

// CloseIssue implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) CloseIssue(ctx context.Context, req *connect.Request[workspacev1.CloseIssueRequest]) (*connect.Response[workspacev1.CloseIssueResponse], error) {
	slog.Info("CloseIssue called", "id", req.Msg.Id)
	issueTask, err := w.IssueUsecase.CloseIssue(ctx, req.Msg.Id)
	if err != nil {
		slog.Error("failed to close issue", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to close issue: %w", err))
	}
	pbIssue, err := toPbIssue(&issueTask)
	if err != nil {
		slog.Error("failed to convert issue to protobuf", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert issue to protobuf: %w", err))
	}
	return connect.NewResponse(&workspacev1.CloseIssueResponse{
		Issue: pbIssue,
	}), nil
}

// CreateIssue implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) CreateIssue(context.Context, *connect.Request[workspacev1.CreateIssueRequest]) (*connect.Response[workspacev1.CreateIssueResponse], error) {
	panic("unimplemented")
}

// CreateTask implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) CreateTask(ctx context.Context, req *connect.Request[workspacev1.CreateTaskRequest]) (*connect.Response[workspacev1.CreateTaskResponse], error) {
	slog.Info("CreateTask called", "title", req.Msg.Title, "description", req.Msg.Description)

	switch req.Msg.Type {
	case workspacev1.TaskType_TASK_TYPE_TODO:
		todoTask, err := w.TodoUsecase.CreateTodo(ctx, req.Msg.Title, req.Msg.Description)
		if err != nil {
			slog.Error("failed to create todo task", "error", err)
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create todo task: %w", err))
		}
		pbTodo := toPbTodo(&todoTask)
		return connect.NewResponse(&workspacev1.CreateTaskResponse{
			Task: &workspacev1.Task{
				Type: &workspacev1.Task_Todo{
					Todo: pbTodo,
				},
			},
		}), nil
	case workspacev1.TaskType_TASK_TYPE_PROGRESS:
		progressTask, err := w.ProgressUsecase.CreateProgress(ctx, req.Msg.Title, req.Msg.Description)
		if err != nil {
			slog.Error("failed to create progress task", "error", err)
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create progress task: %w", err))
		}
		pbProgress, err := toPbProgress(progressTask)
		if err != nil {
			slog.Error("failed to convert progress to protobuf", "error", err)
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert progress to protobuf: %w", err))
		}
		return connect.NewResponse(&workspacev1.CreateTaskResponse{
			Task: &workspacev1.Task{
				Type: &workspacev1.Task_Progress{
					Progress: pbProgress,
				},
			},
		}), nil
	default:
		slog.Error("unknown task type", "type", req.Msg.Type)
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("unknown task type: %s", req.Msg.Type))
	}
}

// DeleteIssue implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) DeleteIssue(context.Context, *connect.Request[workspacev1.DeleteIssueRequest]) (*connect.Response[workspacev1.DeleteIssueResponse], error) {
	panic("unimplemented")
}

// DeleteTask implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) DeleteTask(ctx context.Context, req *connect.Request[workspacev1.DeleteTaskRequest]) (*connect.Response[workspacev1.DeleteTaskResponse], error) {
	slog.Info("DeleteTask called", "id", req.Msg.Id)
	panic("unimplemented")
}

// GetIssue implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) GetIssue(ctx context.Context, req *connect.Request[workspacev1.GetIssueRequest]) (*connect.Response[workspacev1.GetIssueResponse], error) {
	slog.Info("GetIssue called", "id", req.Msg.Id)

	issueTask, ok, err := w.IssueUsecase.GetIssueByID(ctx, req.Msg.Id)
	if err != nil {
		slog.Error("failed to get issue by id", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to get issue by id: %w", err))
	}
	if !ok {
		slog.Error("issue not found", "id", req.Msg.Id)
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("issue with id %s not found", req.Msg.Id))
	}
	pbIssue, err := toPbIssue(issueTask)
	if err != nil {
		slog.Error("failed to convert issue to protobuf", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert issue to protobuf: %w", err))
	}
	return connect.NewResponse(&workspacev1.GetIssueResponse{
		Issue: pbIssue,
	}), nil
}

// GetTask implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) GetTask(context.Context, *connect.Request[workspacev1.GetTaskRequest]) (*connect.Response[workspacev1.GetTaskResponse], error) {
	panic("unimplemented")
}

// InvestigateIssue implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) InvestigateIssue(ctx context.Context, req *connect.Request[workspacev1.InvestigateIssueRequest]) (*connect.Response[workspacev1.InvestigateIssueResponse], error) {
	issueTask, err := w.IssueUsecase.InvestigateIssue(ctx, req.Msg.Id)
	if err != nil {
		slog.Error("failed to investigate issue", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to investigate issue: %w", err))
	}
	pbIssue, err := toPbIssue(issueTask)
	if err != nil {
		slog.Error("failed to convert issue to protobuf", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert issue to protobuf: %w", err))
	}
	return connect.NewResponse(&workspacev1.InvestigateIssueResponse{
		Issue: pbIssue,
	}), nil
}

// ListIssues implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) ListIssues(ctx context.Context, req *connect.Request[workspacev1.ListIssuesRequest]) (*connect.Response[workspacev1.ListIssuesResponse], error) {
	issueTasks, err := w.IssueUsecase.ListIssues(ctx)
	if err != nil {
		slog.Error("failed to list issues", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list issues: %w", err))
	}
	pbIssues := make([]*workspacev1.Issue, len(issueTasks))
	for i, it := range issueTasks {
		pbIssue, err := toPbIssue(&it)
		if err != nil {
			slog.Error("failed to convert issue to protobuf", "error", err)
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert issue to protobuf: %w", err))
		}
		pbIssues[i] = pbIssue
	}
	return connect.NewResponse(&workspacev1.ListIssuesResponse{
		Issues: pbIssues,
	}), nil
}

// ListTasks implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) ListTasks(context.Context, *connect.Request[workspacev1.ListTasksRequest]) (*connect.Response[workspacev1.ListTasksResponse], error) {
	panic("unimplemented")
}

// ResolveIssue implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) ResolveIssue(ctx context.Context, req *connect.Request[workspacev1.ResolveIssueRequest]) (*connect.Response[workspacev1.ResolveIssueResponse], error) {
	issueTask, err := w.IssueUsecase.ResolveIssue(ctx, req.Msg.Id)
	if err != nil {
		slog.Error("failed to resolve issue", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to resolve issue: %w", err))
	}
	pbIssue, err := toPbIssue(issueTask)
	if err != nil {
		slog.Error("failed to convert issue to protobuf", "error", err)
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to convert issue to protobuf: %w", err))
	}
	return connect.NewResponse(&workspacev1.ResolveIssueResponse{
		Issue: pbIssue,
	}), nil
}

// CompleteProgress implements workspacev1connect.WorkspaceServiceHandler.
func (w *WorkSpaceController) CompleteProgress(ctx context.Context, req *connect.Request[workspacev1.CompleteProgressRequest]) (*connect.Response[workspacev1.CompleteProgressResponse], error) {
	progressTask, err := w.ProgressUsecase.CompleteProgress(ctx, req.Msg.Id)
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
	progressTask, err := w.ProgressUsecase.CreateProgress(ctx, req.Msg.Title, req.Msg.Description)
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
	ok, err := w.ProgressUsecase.DeleteProgress(ctx, req.Msg.Id)
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
	progressTask, ok, err := w.ProgressUsecase.GetProgressByID(ctx, req.Msg.Id)
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
	progressTasks, err := w.ProgressUsecase.ListProgress(ctx)
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
	progressTask, err := w.ProgressUsecase.SetSolution(ctx, req.Msg.Id, req.Msg.Solution)
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
	progressTask, err := w.ProgressUsecase.StartProgress(ctx, req.Msg.Id)
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

	switch progressTask.State().Status() {
	case progress.NotStarted:
		pb.State = &workspacev1.Progress_NotStarted{}
	case progress.InProgress:
		if solution, ok := progressTask.State().Solution(); ok {
			pb.State = &workspacev1.Progress_InProgress{
				InProgress: &workspacev1.ProgressInProgressState{
					Solution: &solution,
				},
			}
		} else {
			pb.State = &workspacev1.Progress_InProgress{
				InProgress: &workspacev1.ProgressInProgressState{
					Solution: nil, // No solution provided
				},
			}
		}
	case progress.Completed:
		solution, ok := progressTask.State().Solution()
		if !ok {
			return nil, fmt.Errorf("solution is required for completed state")
		}
		pb.State = &workspacev1.Progress_Completed{
			Completed: &workspacev1.ProgressCompletedState{
				Solution: solution,
			},
		}
	default:
		return nil, fmt.Errorf("unknown progress state: %T", progressTask.State())
	}
	return
}

func toPbIssue(issueTask *issue.IssueTask) (pb *workspacev1.Issue, err error) {
	pb = &workspacev1.Issue{}
	pb.Id = issueTask.Data().ID.String()
	pb.Title = issueTask.Data().Title
	pb.Description = issueTask.Data().Description
	pb.CreatedAt = timestamppb.New(issueTask.Data().CreatedAt)
	pb.UpdatedAt = timestamppb.New(issueTask.Data().UpdatedAt)

	switch issueTask.State().Status() {
	case issue.Open:
		pb.State = &workspacev1.Issue_Open{
			Open: &workspacev1.IssueOpenState{},
		}
	case issue.Investigating:
		if cause, ok := issueTask.State().Cause(); ok {
			pb.State = &workspacev1.Issue_Investigating{
				Investigating: &workspacev1.IssueInvestigatingState{
					Cause: &cause,
				},
			}
		} else {
			pb.State = &workspacev1.Issue_Investigating{
				Investigating: &workspacev1.IssueInvestigatingState{
					Cause: nil, // No cause provided
				},
			}
		}
	case issue.Resolving:
		cause, ok := issueTask.State().Cause()
		if !ok {
			return nil, fmt.Errorf("cause is required for resolving state")
		}
		if solution, ok := issueTask.State().Solution(); ok {
			pb.State = &workspacev1.Issue_Resolution{
				Resolution: &workspacev1.IssueResolvingState{
					Cause:    cause,
					Solution: &solution,
				},
			}
		} else {
			pb.State = &workspacev1.Issue_Resolution{
				Resolution: &workspacev1.IssueResolvingState{
					Cause:    cause,
					Solution: nil,
				},
			}
		}
	case issue.Closed:
		cause, ok := issueTask.State().Cause()
		if !ok {
			return nil, fmt.Errorf("cause is required for closed state")
		}
		solution, ok := issueTask.State().Solution()
		if !ok {
			return nil, fmt.Errorf("solution is required for closed state")
		}
		pb.State = &workspacev1.Issue_Closed{
			Closed: &workspacev1.IssueClosedState{
				Cause:    cause,
				Solution: solution,
			},
		}
	}
	return
}

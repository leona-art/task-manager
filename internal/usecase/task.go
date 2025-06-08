package usecase

type TaskUseCase struct {
	todoUseCase *TodoUseCase
}

func NewTaskUseCase(todoUseCase *TodoUseCase) *TaskUseCase {
	return &TaskUseCase{
		todoUseCase: todoUseCase,
	}
}

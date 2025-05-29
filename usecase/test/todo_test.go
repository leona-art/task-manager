package usecase_test

import (
	"testing"

	"github.com/leona-art/task-manager/domain/entity/todo"
	"github.com/leona-art/task-manager/usecase"
)

type InMemoryRepository struct {
	todos map[todo.TodoId]todo.Todo
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		todos: make(map[todo.TodoId]todo.Todo),
	}
}

// Delete implements todo.TodoRepository.
func (i *InMemoryRepository) Delete(id todo.TodoId) (ok bool, err error) {
	delete(i.todos, id)
	return true, nil
}

// GetByID implements todo.TodoRepository.
func (i *InMemoryRepository) GetByID(id todo.TodoId) (value todo.Todo, ok bool, err error) {
	value, ok = i.todos[id]
	return value, ok, nil
}

// List implements todo.TodoRepository.
func (i *InMemoryRepository) List() (todos []todo.Todo, err error) {
	for _, todo := range i.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

// Save implements todo.TodoRepository.
func (i *InMemoryRepository) Save(todo todo.Todo) error {
	i.todos[todo.ID] = todo
	return nil
}

// Update implements todo.TodoRepository.
func (i *InMemoryRepository) Update(todo todo.Todo) error {
	i.todos[todo.ID] = todo
	return nil
}

func TestTodo(t *testing.T) {
	t.Run("TestTodo", func(t *testing.T) {
		repo := NewInMemoryRepository()
		uc := usecase.NewTodoUseCase(repo)

		todo, err := uc.CreateTodo("Test Title", "Test Description")
		if err != nil {
			t.Fatalf("Failed to create todo: %v", err)
		}
		if todo.IsEmpty() {
			t.Fatal("Created todo should not be empty")
		}

		if len(repo.todos) != 1 {
			t.Fatalf("Expected 1 todo in repository, got %d", len(repo.todos))
		}
	})

	t.Run("TestGetTodoByID", func(t *testing.T) {
		repo := NewInMemoryRepository()
		uc := usecase.NewTodoUseCase(repo)

		todo, err := uc.CreateTodo("Test Title", "Test Description")
		if err != nil {
			t.Fatalf("Failed to create todo: %v", err)
		}

		gotTodo, ok, err := uc.GetTodoByID(todo.ID)
		if err != nil {
			t.Fatalf("Failed to get todo by ID: %v", err)
		}
		if !ok {
			t.Fatal("Expected todo to be found")
		}
		if gotTodo.ID != todo.ID {
			t.Fatalf("Expected todo ID %s, got %s", todo.ID, gotTodo.ID)
		}
	})

	t.Run("TestUpdateTodo", func(t *testing.T) {
		repo := NewInMemoryRepository()
		uc := usecase.NewTodoUseCase(repo)

		beforeTitle := "Test Title"
		afterTitle := "Updated Title"
		todo, err := uc.CreateTodo(beforeTitle, "Test Description")
		if err != nil {
			t.Fatalf("Failed to create todo: %v", err)
		}
		todo.Info.Title = afterTitle
		if err := uc.UpdateTodo(todo); err != nil {
			t.Fatalf("Failed to update todo: %v", err)
		}
		updatedTodo, ok, err := uc.GetTodoByID(todo.ID)
		if err != nil {
			t.Fatalf("Failed to get updated todo: %v", err)
		}
		if !ok {
			t.Fatal("Expected updated todo to be found")
		}
		if updatedTodo.Info.Title != afterTitle {
			t.Fatalf("Expected todo title to be '%s', got '%s'", afterTitle, updatedTodo.Info.Title)
		}

	})
}

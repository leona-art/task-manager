package task

import "testing"

func TestStatus(t *testing.T) {
	t.Run("New Pending Status", func(t *testing.T) {
		status := NewTaskPendingState()
		if status.Status() != Pending {
			t.Errorf("Expected status to be '%s', got '%s'", Pending, status.Status())
		}
		if _, ok := status.Resolution(); ok {
			t.Error("Expected no resolution for pending status")
		}
	})

	t.Run("New InProgress Status", func(t *testing.T) {
		status := NewTaskInProgressState()
		if status.Status() != InProgress {
			t.Errorf("Expected status to be '%s', got '%s'", InProgress, status.Status())
		}
		if _, ok := status.Resolution(); ok {
			t.Error("Expected no resolution for in-progress status")
		}
	})

	t.Run("New InProgress Status with Resolution", func(t *testing.T) {
		resolution := "Task is being worked on"
		status := NewTaskInProgressStateWithResolution(resolution)
		if status.Status() != InProgress {
			t.Errorf("Expected status to be '%s', got '%s'", InProgress, status.Status())
		}
		if res, ok := status.Resolution(); !ok || res != resolution {
			t.Errorf("Expected resolution to be '%s', got '%s'", resolution, res)
		}
	})

	t.Run("New Completed Status", func(t *testing.T) {
		resolution := "Task completed successfully"
		status := NewTaskCompletedState(resolution)
		if status.Status() != Completed {
			t.Errorf("Expected status to be '%s', got '%s'", Completed, status.Status())
		}
		if res, ok := status.Resolution(); !ok || res != resolution {
			t.Errorf("Expected resolution to be '%s', got '%s'", resolution, res)
		}
	})

	t.Run("Pending To InProgress Transition", func(t *testing.T) {
		pendingStatus := NewTaskPendingState()
		candidate := pendingStatus.Candidate()
		if progress, ok := candidate[InProgress]; ok {
			inProgressStatus := progress()
			if inProgressStatus.Status() != InProgress {
				t.Errorf("Expected status to be '%s', got '%s'", InProgress, inProgressStatus.Status())
			}
			if _, ok := inProgressStatus.Resolution(); ok {
				t.Error("Expected no resolution for in-progress status")
			}
		} else {
			t.Error("Expected InProgress candidate status not found")
		}
	})

	t.Run("InProgress To Completed Transition", func(t *testing.T) {
		inProgressStatus := NewTaskInProgressStateWithResolution("Task is being worked on")
		candidate := inProgressStatus.Candidate()
		if complete, ok := candidate[Completed]; ok {
			completedStatus := complete()
			if completedStatus.Status() != Completed {
				t.Errorf("Expected status to be '%s', got '%s'", Completed, completedStatus.Status())
			}
			if _, ok := completedStatus.Resolution(); !ok {
				t.Error("Expected no resolution for completed status")
			}
		} else {
			t.Error("Expected Completed candidate status not found")
		}
	})

	t.Run("Failed Transition Not resolution", func(t *testing.T) {
		inProgressStatus := NewTaskInProgressState()
		if _, ok := inProgressStatus.Candidate()[Completed]; ok {
			t.Error("Expected no Completed candidate status for in-progress without resolution")
		}
	})
}

func TestTask(t *testing.T) {
	t.Run("New task", func(t *testing.T) {
		task, err := NewTask("Test Task", "This is a test task")
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
		if task.State.Status() != Pending {
			t.Errorf("Expected task status to be '%s', got '%s'", Pending, task.State.Status())
		}
	})

	t.Run("Start task", func(t *testing.T) {
		task, err := NewTask("Test Task", "This is a test task")
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
		err = task.Start()
		if err != nil {
			t.Errorf("Expected no error, got '%v'", err)
		}
		if task.State.Status() != InProgress {
			t.Errorf("Expected task status to be '%s', got '%s'", InProgress, task.State.Status())
		}
	})

}

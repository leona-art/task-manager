package task_info

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewTaskInfo(t *testing.T) {
	title := "Test Task"
	description := "This is a test task description."

	taskInfo := NewTaskInfo(title, description)

	if taskInfo.Title != title {
		t.Errorf("expected title %s, got %s", title, taskInfo.Title)
	}
	if taskInfo.Description != description {
		t.Errorf("expected description %s, got %s", description, taskInfo.Description)
	}
	if taskInfo.ID == "" {
		t.Error("expected ID to be generated, but it was empty")
	}
	if taskInfo.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set, but it was zero")
	}
	_, err := uuid.Parse(taskInfo.ID)
	if err != nil {
		t.Errorf("failed to parse UUID: %v", err)
	}
}

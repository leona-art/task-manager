package taskinfo

import (
	"testing"
)

func TestNewTaskInfo(t *testing.T) {
	title := "Test Task"
	description := "This is a test task description."

	taskInfo := NewTaskInfo(title, description)

	if taskInfo.IsEmpty() {
		t.Error("Expected task info to be not empty, but it is empty")
	}
}

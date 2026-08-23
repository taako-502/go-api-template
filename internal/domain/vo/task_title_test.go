package vo

import (
	"errors"
	"testing"
)

func TestNewTaskTitle(t *testing.T) {
	title, err := NewTaskTitle("  write tests  ")
	if err != nil {
		t.Fatal(err)
	}
	if title.String() != "write tests" {
		t.Fatalf("title=%q", title)
	}
}

func TestNewTaskTitleRejectsEmptyValue(t *testing.T) {
	if _, err := NewTaskTitle("   "); !errors.Is(err, ErrInvalidTaskTitle) {
		t.Fatalf("error=%v", err)
	}
}

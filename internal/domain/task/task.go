package task

import (
	"errors"
	"strings"
	"time"
	"uuid"
)

var (
	ErrNotFound     = errors.New("task not found")
	ErrInvalidTitle = errors.New("title is required")
)

type Task struct {
	ID        uuid.UUID
	Title     string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(title string, now time.Time) (Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, ErrInvalidTitle
	}
	now = now.UTC()
	return Task{ID: uuid.NewV7(), Title: title, CreatedAt: now, UpdatedAt: now}, nil
}
func (t *Task) Update(title string, done bool, now time.Time) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return ErrInvalidTitle
	}
	t.Title, t.Done, t.UpdatedAt = title, done, now.UTC()
	return nil
}

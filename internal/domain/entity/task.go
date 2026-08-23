package entity

import (
	"time"
	"uuid"

	"github.com/taako-502/go-api-template/internal/domain/vo"
)

type Task struct {
	ID        uuid.UUID
	Title     vo.TaskTitle
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTask(title vo.TaskTitle, now time.Time) Task {
	now = now.UTC()
	return Task{ID: uuid.NewV7(), Title: title, CreatedAt: now, UpdatedAt: now}
}

func (t *Task) Update(title vo.TaskTitle, done bool, now time.Time) {
	t.Title, t.Done, t.UpdatedAt = title, done, now.UTC()
}

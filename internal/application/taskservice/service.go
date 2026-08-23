package taskservice

import (
	"context"
	"github.com/taako-502/go-api-template/internal/application/page"
	"github.com/taako-502/go-api-template/internal/domain/task"
	"time"
	"uuid"
)

type Service struct {
	repository task.Repository
	now        func() time.Time
}

func New(repository task.Repository) Service { return Service{repository: repository, now: time.Now} }
func (s Service) Create(ctx context.Context, title string) (task.Task, error) {
	value, err := task.New(title, s.now())
	if err != nil {
		return task.Task{}, err
	}
	if err := s.repository.Create(ctx, value); err != nil {
		return task.Task{}, err
	}
	return value, nil
}
func (s Service) List(ctx context.Context) (page.Page[task.Task], error) {
	items, err := s.repository.List(ctx)
	return page.Page[task.Task]{Items: items, Total: len(items)}, err
}
func (s Service) Get(ctx context.Context, id uuid.UUID) (task.Task, error) {
	return s.repository.Find(ctx, id)
}
func (s Service) Update(ctx context.Context, id uuid.UUID, title string, done bool) (task.Task, error) {
	value, err := s.repository.Find(ctx, id)
	if err != nil {
		return task.Task{}, err
	}
	if err := value.Update(title, done, s.now()); err != nil {
		return task.Task{}, err
	}
	if err := s.repository.Update(ctx, value); err != nil {
		return task.Task{}, err
	}
	return value, nil
}
func (s Service) Delete(ctx context.Context, id uuid.UUID) error { return s.repository.Delete(ctx, id) }

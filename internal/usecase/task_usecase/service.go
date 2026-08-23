package task_usecase

import (
	"context"
	"time"
	"uuid"

	"github.com/taako-502/go-api-template/internal/domain/entity"
	"github.com/taako-502/go-api-template/internal/domain/repository"
	"github.com/taako-502/go-api-template/internal/domain/vo"
	"github.com/taako-502/go-api-template/internal/usecase/page"
)

type Service struct {
	repository repository.TaskRepository
	now        func() time.Time
}

func New(repository repository.TaskRepository) Service {
	return Service{repository: repository, now: time.Now}
}

func (s Service) Create(ctx context.Context, title string) (entity.Task, error) {
	taskTitle, err := vo.NewTaskTitle(title)
	if err != nil {
		return entity.Task{}, err
	}
	value := entity.NewTask(taskTitle, s.now())
	if err := s.repository.Create(ctx, value); err != nil {
		return entity.Task{}, err
	}
	return value, nil
}

func (s Service) List(ctx context.Context) (page.Page[entity.Task], error) {
	items, err := s.repository.List(ctx)
	return page.Page[entity.Task]{Items: items, Total: len(items)}, err
}

func (s Service) Get(ctx context.Context, id uuid.UUID) (entity.Task, error) {
	return s.repository.Find(ctx, id)
}

func (s Service) Update(ctx context.Context, id uuid.UUID, title string, done bool) (entity.Task, error) {
	taskTitle, err := vo.NewTaskTitle(title)
	if err != nil {
		return entity.Task{}, err
	}
	value, err := s.repository.Find(ctx, id)
	if err != nil {
		return entity.Task{}, err
	}
	value.Update(taskTitle, done, s.now())
	if err := s.repository.Update(ctx, value); err != nil {
		return entity.Task{}, err
	}
	return value, nil
}

func (s Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}

package repository

import (
	"context"
	"errors"
	"uuid"

	"github.com/taako-502/go-api-template/internal/domain/entity"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository interface {
	Create(context.Context, entity.Task) error
	List(context.Context) ([]entity.Task, error)
	Find(context.Context, uuid.UUID) (entity.Task, error)
	Update(context.Context, entity.Task) error
	Delete(context.Context, uuid.UUID) error
}

package task

import (
	"context"
	"uuid"
)

type Repository interface {
	Create(context.Context, Task) error
	List(context.Context) ([]Task, error)
	Find(context.Context, uuid.UUID) (Task, error)
	Update(context.Context, Task) error
	Delete(context.Context, uuid.UUID) error
}

package memory

import (
	"context"
	"github.com/taako-502/go-api-template/internal/domain/task"
	"slices"
	"sync"
	"uuid"
)

type TaskRepository struct {
	mu    sync.RWMutex
	items map[uuid.UUID]task.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{items: make(map[uuid.UUID]task.Task)}
}
func (r *TaskRepository) Create(_ context.Context, value task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[value.ID] = value
	return nil
}
func (r *TaskRepository) List(_ context.Context) ([]task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := make([]task.Task, 0, len(r.items))
	for _, v := range r.items {
		items = append(items, v)
	}
	slices.SortFunc(items, func(a, b task.Task) int { return a.ID.Compare(b.ID) })
	return items, nil
}
func (r *TaskRepository) Find(_ context.Context, id uuid.UUID) (task.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.items[id]
	if !ok {
		return task.Task{}, task.ErrNotFound
	}
	return v, nil
}
func (r *TaskRepository) Update(_ context.Context, v task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[v.ID]; !ok {
		return task.ErrNotFound
	}
	r.items[v.ID] = v
	return nil
}
func (r *TaskRepository) Delete(_ context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return task.ErrNotFound
	}
	delete(r.items, id)
	return nil
}

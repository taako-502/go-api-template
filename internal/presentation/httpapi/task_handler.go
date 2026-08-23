package httpapi

import (
	"context"
	"errors"
	"github.com/danielgtaylor/huma/v2"
	"github.com/taako-502/go-api-template/internal/application/taskservice"
	"github.com/taako-502/go-api-template/internal/domain/task"
	"time"
	"uuid"
)

type TaskHandler struct{ service taskservice.Service }

func NewTaskHandler(service taskservice.Service) TaskHandler { return TaskHandler{service: service} }

type taskBody struct {
	ID        string    `json:"id" format:"uuid"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type createTaskInput struct {
	Body struct {
		Title string `json:"title" minLength:"1" maxLength:"200"`
	}
}
type taskOutput struct{ Body taskBody }
type listTasksOutput struct {
	Body struct {
		Items []taskBody `json:"items"`
		Total int        `json:"total"`
	}
}
type taskIDInput struct {
	ID uuid.UUID `path:"id" format:"uuid"`
}
type updateTaskInput struct {
	ID   uuid.UUID `path:"id" format:"uuid"`
	Body struct {
		Title string `json:"title" minLength:"1" maxLength:"200"`
		Done  bool   `json:"done"`
	}
}
type emptyOutput struct{}

func (h TaskHandler) Create(ctx context.Context, in *createTaskInput) (*taskOutput, error) {
	v, e := h.service.Create(ctx, in.Body.Title)
	if e != nil {
		return nil, mapError(e)
	}
	return &taskOutput{Body: toTaskBody(v)}, nil
}
func (h TaskHandler) List(ctx context.Context, _ *struct{}) (*listTasksOutput, error) {
	v, e := h.service.List(ctx)
	if e != nil {
		return nil, mapError(e)
	}
	m := v.Map(toTaskBody)
	out := &listTasksOutput{}
	out.Body.Items, out.Body.Total = m.Items, m.Total
	return out, nil
}
func (h TaskHandler) Get(ctx context.Context, in *taskIDInput) (*taskOutput, error) {
	v, e := h.service.Get(ctx, in.ID)
	if e != nil {
		return nil, mapError(e)
	}
	return &taskOutput{Body: toTaskBody(v)}, nil
}
func (h TaskHandler) Update(ctx context.Context, in *updateTaskInput) (*taskOutput, error) {
	v, e := h.service.Update(ctx, in.ID, in.Body.Title, in.Body.Done)
	if e != nil {
		return nil, mapError(e)
	}
	return &taskOutput{Body: toTaskBody(v)}, nil
}
func (h TaskHandler) Delete(ctx context.Context, in *taskIDInput) (*emptyOutput, error) {
	if e := h.service.Delete(ctx, in.ID); e != nil {
		return nil, mapError(e)
	}
	return &emptyOutput{}, nil
}
func toTaskBody(v task.Task) taskBody {
	return taskBody{ID: v.ID.String(), Title: v.Title, Done: v.Done, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt}
}
func mapError(err error) error {
	switch {
	case errors.Is(err, task.ErrNotFound):
		return huma.Error404NotFound("task not found")
	case errors.Is(err, task.ErrInvalidTitle):
		return huma.Error400BadRequest(err.Error())
	default:
		return huma.Error500InternalServerError("internal server error")
	}
}

package usecase

import (
	"awesomeProject/internal/domain"
	"context"
)

type TaskUseCase interface {
	FindAll(ctx context.Context) ([]domain.Task, error)
	FindByID(ctx context.Context, id uint) (domain.Task, error)
	Save(ctx context.Context, user domain.Task) (domain.Task, error)
	Delete(ctx context.Context, user domain.Task) error
	Update(ctx context.Context, task domain.Task) interface{}
	Post(ctx context.Context, task domain.Task) (interface{}, interface{})
}

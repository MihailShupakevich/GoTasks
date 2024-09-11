package repository

import (
	"awesomeProject/internal/domain"
	"context"
)

type TaskRepository interface {
	FindAll(ctx context.Context) ([]domain.Task, error)
	FindByID(ctx context.Context, id uint) (domain.Task, error)
	Save(ctx context.Context, user domain.Task) (domain.Task, error)
	Delete(ctx context.Context, user domain.Task) error
}

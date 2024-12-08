package interfaces

import (
	"awesomeProject/internal/domain"
	_ "awesomeProject/internal/usecase/task"
	interfaces "awesomeProject/internal/usecase/task"
	"context"
)

type taskUseCase struct {
	taskRepo interfaces.TaskUseCase
}

func (c *taskUseCase) FindAll(ctx context.Context) ([]domain.Task, error) {
	users, err := c.taskRepo.FindAll(ctx)
	return users, err
}

func (c *taskUseCase) FindByID(ctx context.Context, id uint) (domain.Task, error) {
	user, err := c.taskRepo.FindByID(ctx, id)
	return user, err
}

func (c *taskUseCase) Save(ctx context.Context, user domain.Task) (domain.Task, error) {
	user, err := c.taskRepo.Save(ctx, user)

	return user, err
}

func (c *taskUseCase) Delete(ctx context.Context, user domain.Task) error {
	err := c.taskRepo.Delete(ctx, user)

	return err
}

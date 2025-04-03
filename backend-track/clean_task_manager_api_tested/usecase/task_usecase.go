package usecase

import (
	"clean_task_manager_api_tested/domain"
	"context"
	"time"
)

type TaskUsecase struct {
	taskRepositoryInteface domain.TaskRepositoryInteface
	contextTimeout         time.Duration
}

// lets implement the TaskUsecase interface methods

func (tu *TaskUsecase) Create(c context.Context, task *domain.Task) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepositoryInteface.Create(ctx, task)
}

func (tu *TaskUsecase) GetAllTasks(c context.Context) ([]domain.Task, error) {

	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepositoryInteface.GetAllTasks(ctx)
}

func (tu *TaskUsecase) GetTaskByID(c context.Context, taskID string) (*domain.Task, error) {

	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepositoryInteface.GetTaskByID(ctx, taskID)
}

func (tu *TaskUsecase) UpdateTask(c context.Context, taskID string, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepositoryInteface.UpdateTask(ctx, taskID, task)
}

func (tu *TaskUsecase) DeleteTask(c context.Context, taskID string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepositoryInteface.DeleteTask(ctx, taskID)
}

func NewTaskUsecase(taskRepositoryInterface domain.TaskRepositoryInteface, timeout time.Duration) domain.TaskUsecaseInteface {

	return &TaskUsecase{
		taskRepositoryInteface: taskRepositoryInterface,
		contextTimeout:         timeout,
	}
}

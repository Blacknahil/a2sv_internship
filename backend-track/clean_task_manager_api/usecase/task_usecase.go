package usecase

import (
	"clean-task-manager-api/domain"
	"context"
	"time"
)

type taskUsecase struct {
	taskRepositoryInteface domain.TaskRepositoryInteface
	contextTimeout         time.Duration
}

// lets implement the taskUsecase interface methods

func (tu *taskUsecase) Create(c context.Context, task *domain.Task) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepositoryInteface.Create(ctx, task)
}

func (tu *taskUsecase) GetAllTasks(c context.Context) ([]domain.Task, error) {

	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepositoryInteface.GetAllTasks(ctx)
}

func (tu *taskUsecase) GetTaskByID(c context.Context, taskID string) (*domain.Task, error) {

	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepositoryInteface.GetTaskByID(ctx, taskID)
}

func (tu *taskUsecase) UpdateTask(c context.Context, taskID string, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepositoryInteface.UpdateTask(ctx, taskID, task)
}

func (tu *taskUsecase) DeleteTask(c context.Context, taskID string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepositoryInteface.DeleteTask(ctx, taskID)
}

func NewTaskUsecase(taskRepositoryInterface domain.TaskRepositoryInteface, timeout time.Duration) domain.TaskUsecaseInteface {

	return &taskUsecase{
		taskRepositoryInteface: taskRepositoryInterface,
		contextTimeout:         timeout,
	}
}

// kfkjsdfkj
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//dkjjkjsdkfjkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf dfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
// jkdnfkfjsfdkjskdj

// bbfkjdsbkj
package domain

import "context"

type TaskRepositoryInteface interface {
	Create(c context.Context, task *Task) (*Task, error)
	GetAllTasks(c context.Context) ([]Task, error)
	GetTaskByID(c context.Context, taskID string) (*Task, error)
	UpdateTask(c context.Context, taskID string, task *Task) error
	DeleteTask(c context.Context, taskID string) error
}

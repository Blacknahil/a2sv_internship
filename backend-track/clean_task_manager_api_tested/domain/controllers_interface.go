package domain

import "github.com/gin-gonic/gin"

// UserControllerInterface defines the methods for the UserController
type UserControllerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Promote(c *gin.Context)
}

// TaskControllerInterface defines the methods for the TaskController
type TaskControllerInterface interface {
	Create(c *gin.Context)
	GetAllTasks(c *gin.Context)
	GetTasksById(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

package router

import (
	controllers "task_manager_api_auth/controller"
	"task_manager_api_auth/data"
	"task_manager_api_auth/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(r *gin.Engine) {
	taskCollection := data.TaskCollection
	userCollection := data.UserCollection

	userServices := data.UserServices{UserCollection: userCollection}
	taskServices := data.TaskServices{TaskCollection: taskCollection}

	taskController := &controllers.TaskController{TaskServices: taskServices}
	userController := &controllers.UserController{UserServices: userServices}

	user := r.Group("/users")
	{
		user.GET("/", middleware.AuthMiddleware(), middleware.AdminMiddleware(), userController.GetAllUsers)
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}

	task := r.Group("/tasks")
	{
		task.POST("/", middleware.AuthMiddleware(), middleware.AdminMiddleware(), taskController.CreateTask)
		task.GET("/", middleware.AuthMiddleware(), taskController.GetAllTasks)
		task.GET("/:id", middleware.AuthMiddleware(), taskController.GetTasksById)
		task.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), taskController.UpdateTask)
		task.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), taskController.DeleteTask)
	}

}

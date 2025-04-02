package router

import (
	"clean-task-manager-api/delivery/controllers"
	"clean-task-manager-api/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetUp(r *gin.Engine, userController *controllers.UserController, taskController *controllers.TaskController) {

	authMiddleware := infrastructure.AuthMiddleware()
	adminMiddleware := infrastructure.AdminMiddleware()

	user := r.Group("/users")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
		user.POST("/promote/:id", authMiddleware, adminMiddleware, userController.Promote)
	}

	task := r.Group("/tasks")
	{
		task.POST("/", authMiddleware, adminMiddleware, taskController.Create)
		task.GET("/", authMiddleware, taskController.GetAllTasks)
		task.GET("/:id", authMiddleware, taskController.GetTasksById)
		task.PUT("/:id", authMiddleware, adminMiddleware, taskController.UpdateTask)
		task.DELETE("/:id", authMiddleware, adminMiddleware, taskController.DeleteTask)
	}
}

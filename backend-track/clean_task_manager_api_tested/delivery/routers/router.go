package router

import (
	"clean_task_manager_api_tested/domain"
	"clean_task_manager_api_tested/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetUp(r *gin.Engine, userController domain.UserControllerInterface, taskController domain.TaskControllerInterface) {

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

package router

import (
	controllers "task_manager_api_auth/controller"
	"task_manager_api_auth/data"

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
		user.GET("/", userController.GetAllUsers)
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}

	task := r.Group("/tasks")
	{
		task.POST("/", taskController.CreateTask)
		task.GET("/", taskController.GetAllTasks)
		task.GET("/:id", taskController.GetTasksById)
		task.PUT("/:id", taskController.UpdateTask)
		task.DELETE("/:id", taskController.DeleteTask)
	}

}

//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
// kjdjkdskjksdj

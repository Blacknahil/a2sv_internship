package router

import (
	"task_manager_api/controllers"
	"task_manager_api/data"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpRouter(client *mongo.Client) *gin.Engine {

	taskCollection := client.Database("tasksDB").Collection("tasks")
	taskServices := &data.TaskService{Collection: *taskCollection}
	taskController := &controllers.TaskController{Services: taskServices}

	router := gin.Default()

	router.GET("/tasks", taskController.GetAllTasks)
	router.GET("/tasks/:id", taskController.GetTasksById)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)
	router.POST("/tasks", taskController.CreateTask)

	return router
}

//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk

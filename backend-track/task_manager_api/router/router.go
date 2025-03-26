// jkdfjgjksdjbfjksdbfjsdjjhsdfhjshjddhjsdjhjksdjk
package router

import (
	"task_manager_api/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:id", controllers.GetTasksById)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.UpdateTask)
	router.POST("/tasks", controllers.CreateTask)

	// router.Run("localhost:8080")
	// router.Run()
	// fmt.Println("Task manager API")
	return router
}

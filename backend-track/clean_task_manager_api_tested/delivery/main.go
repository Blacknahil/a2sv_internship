// jdhjfdhjhj
package main

import (
	"clean_task_manager_api_tested/delivery/controllers"
	router "clean_task_manager_api_tested/delivery/routers"
	"clean_task_manager_api_tested/delivery/utils"
	repository "clean_task_manager_api_tested/repositories"
	"clean_task_manager_api_tested/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// intialize database
	database := utils.IntializeMongoBD()

	// intalize repository imple
	userRepo := repository.NewUserRepositoryImpl(database, "users")
	taskRepo := repository.NewTaskRepositoryImpl(database, "tasks")

	// initialize usecase
	userUsecase := usecase.NewUserUsecase(userRepo, 10*time.Second)
	taskUsecase := usecase.NewTaskUsecase(taskRepo, 10*time.Second)

	// initialize controller
	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	router.SetUp(r, userController, taskController)
	r.Run(":8080")

}

// djhbsdjhjhsdh
// kjjksjkjks
// / djh jsjdhjh

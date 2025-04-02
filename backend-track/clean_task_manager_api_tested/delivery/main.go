// jdhjfdhjhj
package main

import (
	"clean-task-manager-api/delivery/controllers"
	router "clean-task-manager-api/delivery/routers"
	"clean-task-manager-api/delivery/utils"
	repository "clean-task-manager-api/repositories"
	"clean-task-manager-api/usecase"
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
	userUsecase := usecase.NewUserUseCase(userRepo, 10*time.Second)
	taskUsecase := usecase.NewTaskUsecase(taskRepo, 10*time.Second)

	// initialize controller
	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	router.SetUp(r, userController, taskController)
	r.Run(":8080")

}

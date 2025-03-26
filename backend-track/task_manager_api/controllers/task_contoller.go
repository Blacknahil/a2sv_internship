package controllers

import (
	"net/http"
	"task_manager_api/data"
	"task_manager_api/models"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"tasks": data.GetAllTasks()})

}

func GetTasksById(ctx *gin.Context) {

	id := ctx.Param("id")
	task, err := data.GetTaskById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func UpdateTask(ctx *gin.Context) {

	id := ctx.Param("id")

	var updatedTask models.Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.UpdateTask(updatedTask, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
		return

	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task updated "})

}

func DeleteTask(ctx *gin.Context) {

	id := ctx.Param("id")

	err := data.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})

}

func CreateTask(ctx *gin.Context) {

	var newTask models.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.CreateTask(newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkddjkjksdkjskjdjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsdjjhsdfhjshjddhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjdhjshdjhjshjhjshjshjhj
// sdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//

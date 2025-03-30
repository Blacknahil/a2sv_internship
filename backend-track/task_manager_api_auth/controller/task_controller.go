package controllers

import (
	"context"
	"net/http"
	"task_manager_api_auth/data"
	"task_manager_api_auth/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskServices data.TaskServices
}

func (tc *TaskController) CreateTask(ctx *gin.Context) {

	var newTask models.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// // fmt.Println("controller", newTask)
	// // gin.Logger(newTask)

	createdTask, err := tc.TaskServices.CreateTask(context.Background(), &newTask)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully created task", "task": createdTask})
}

func (tc *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := tc.TaskServices.GetAllTasks(context.Background())

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})

}

func (tc *TaskController) GetTasksById(ctx *gin.Context) {

	id := ctx.Param("id")
	task, err := tc.TaskServices.GetTaskById(context.Background(), id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"task": task})
}

func (tc *TaskController) UpdateTask(ctx *gin.Context) {

	id := ctx.Param("id")

	var updatedTask models.Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := tc.TaskServices.UpdateTask(context.Background(), id, updatedTask)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task Updated successfully"})
}

func (tc *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id") // Get the task ID from the URL parameter

	err := tc.TaskServices.DeleteTask(context.Background(), id)
	if err != nil {
		if err.Error() == "task not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}

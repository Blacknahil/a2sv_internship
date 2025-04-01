package controllers

import (
	"clean-task-manager-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecaseInteface
}

func NewTaskController(tu domain.TaskUsecaseInteface) *TaskController {
	return &TaskController{TaskUsecase: tu}
}

func (tc *TaskController) Create(c *gin.Context) {

	var newTask domain.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := tc.TaskUsecase.Create(c, &newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created task", "task": createdTask})
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.TaskUsecase.GetAllTasks(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})

}

func (tc *TaskController) GetTasksById(c *gin.Context) {

	id := c.Param("id")
	task, err := tc.TaskUsecase.GetTaskByID(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {

	id := c.Param("id")

	var updatedTask domain.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := tc.TaskUsecase.UpdateTask(c, id, &updatedTask)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Task Updated successfully"})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id") // Get the task ID from the URL parameter

	err := tc.TaskUsecase.DeleteTask(c, id)
	if err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}

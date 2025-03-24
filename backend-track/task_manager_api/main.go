package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

// Mock data for tasks
var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func main() {

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200,
			gin.H{"message": "pong from nahom"})
	})

	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", getTaskById)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.POST("/tasks", createTask)

	// router.Run("localhost:8080")
	router.Run()
	fmt.Println("Task manager API")
}

func getAllTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func getTaskById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, task := range tasks {

		if task.ID == id {
			ctx.JSON(http.StatusOK, task)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not founc"})
}

func updateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {

		if task.ID == id {

			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}

			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}

			ctx.JSON(http.StatusOK, gin.H{"message": "Task Updated"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})

}

func deleteTask(ctx *gin.Context) {

	id := ctx.Param("id")

	for i, task := range tasks {

		if task.ID == id {

			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"messaage": "Task removed"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func createTask(ctx *gin.Context) {

	var newTask Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task Created"})

}

// djbskjdbkjds
// dfhjshjdjhsd
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk

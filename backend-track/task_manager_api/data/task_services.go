package data

import (
	"errors"
	"task_manager_api/models"
	"time"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func GetAllTasks() []models.Task {
	return tasks
}

func GetTaskById(id string) (*models.Task, error) {

	for _, task := range tasks {

		if task.ID == id {
			// ctx.JSON(http.StatusOK, task)
			return &task, nil
		}
	}

	// ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not founc"})
	return &models.Task{}, errors.New("task not found")
}

func UpdateTask(updatedTask models.Task, id string) error {

	for i, task := range tasks {

		if task.ID == id {

			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}

			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}

			// ctx.JSON(http.StatusOK, gin.H{"message": "Task Updated"})
			return nil
		}
	}

	// ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
	return errors.New("Task not found")

}

func DeleteTask(id string) error {

	for i, task := range tasks {

		if task.ID == id {

			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}

	return errors.New("Task not found")

}

func CreateTask(newTask models.Task) {

	tasks = append(tasks, newTask)
	return

}

// djbskjdbkjds
// dfhjshjdjhsd
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
// fdjh jhfsdjhsjhdjh/jkdfbjkgkdfbjdfdjhsdhjhjsdhjjkd
// dhjfsdjhhj
// dhjsdjhjhsdjhhfhjdfj

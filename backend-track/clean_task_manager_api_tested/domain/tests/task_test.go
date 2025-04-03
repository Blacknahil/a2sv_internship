package domain_test

import (
	"clean_task_manager_api_tested/domain"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskMarshalJSON(t *testing.T) {
	// Arrange: Create a sample task
	objectID := primitive.NewObjectID()
	dueDate := time.Now()
	task := domain.Task{
		ID:          objectID,
		Title:       "Sample Task",
		Description: "This is a sample task",
		DueDate:     dueDate,
		Status:      "Pending",
	}

	// Act: Marshal the task to JSON
	jsonData, err := json.Marshal(task)

	// Assert: Check for no errors
	assert.NoError(t, err)

	// Assert: Check the JSON output
	expectedJSON := `{
        "id": "` + objectID.Hex() + `",
        "title": "Sample Task",
        "description": "This is a sample task",
        "due_date": "` + dueDate.Format(time.RFC3339) + `",
        "status": "Pending"
    }`
	assert.JSONEq(t, expectedJSON, string(jsonData))
}

func TestTaskMarshalJSON_EmptyID(t *testing.T) {
	// Arrange: Create a task with an empty ID
	dueDate := time.Now().Truncate(time.Millisecond) // Truncate to milliseconds for consistency
	task := domain.Task{
		Title:       "Sample Task",
		Description: "This is a sample task",
		DueDate:     dueDate,
		Status:      "Pending",
	}

	// Act: Marshal the task to JSON
	jsonData, err := json.Marshal(task)

	// Assert: Check for no errors
	assert.NoError(t, err)

	// Format the DueDate to exclude milliseconds for the expected JSON
	expectedDueDate := dueDate.Format("2006-01-02T15:04:05-07:00")

	// Assert: Check the JSON output
	expectedJSON := `{
        "id": "000000000000000000000000",
        "title": "Sample Task",
        "description": "This is a sample task",
        "due_date": "` + expectedDueDate + `",
        "status": "Pending"
    }`
	assert.JSONEq(t, expectedJSON, string(jsonData))
}

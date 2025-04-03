package controllers_test

import (
	"bytes"
	"clean_task_manager_api_tested/delivery/controllers"
	"clean_task_manager_api_tested/domain"
	"clean_task_manager_api_tested/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskControllerTestSuite defines the test suite
type TaskControllerTestSuite struct {
	suite.Suite
	mockUsecase *mocks.TaskUsecaseInteface
	controller  *controllers.TaskController
}

// SetupTest initializes the test setup
func (suite *TaskControllerTestSuite) SetupTest() {
	suite.mockUsecase = new(mocks.TaskUsecaseInteface)
	suite.controller = controllers.NewTaskController(suite.mockUsecase)
}

// TestCreateTask_Success tests task creation with valid data
func (suite *TaskControllerTestSuite) TestCreateTask_Success() {
	// Arrange: Mock Data
	mockTask := domain.Task{
		Title:       "New Task",
		Description: "Task Description",
	}

	mockCreatedTask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       mockTask.Title,
		Description: mockTask.Description,
	}

	mockBody, _ := json.Marshal(mockTask)

	// Expectation: `Create` should be called once and return the mock response
	suite.mockUsecase.On("Create", mock.Anything, &mockTask).Return(&mockCreatedTask, nil).Once()

	// Create a fake gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(mockBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act: Call the controller
	suite.controller.Create(c)

	// Assert: Check response
	expectedResponse, _ := json.Marshal(gin.H{
		"message": "Successfully created task",
		"task":    mockCreatedTask,
	})

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	assert.JSONEq(suite.T(), string(expectedResponse), w.Body.String())

	// Verify that the expectation was met
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestCreateTask_InvalidPayload tests task creation with invalid JSON
func (suite *TaskControllerTestSuite) TestCreateTask_InvalidPayload() {
	// Arrange: Invalid JSON (missing required fields)
	invalidBody := `{"title":""}` // Missing "description"

	// Create a fake gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(invalidBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act: Call the controller
	suite.controller.Create(c)

	// Assert: Expect 400 Bad Request
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

// TestCreateTask_Failure tests task creation failure due to internal error
func (suite *TaskControllerTestSuite) TestCreateTask_Failure() {
	// Arrange: Mock Data
	mockTask := domain.Task{
		Title:       "New Task",
		Description: "Task Description",
	}

	mockBody, _ := json.Marshal(mockTask)

	// Expectation: `Create` should return an error
	suite.mockUsecase.On("Create", mock.Anything, &mockTask).Return(nil, assert.AnError).Once()

	// Create a fake gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(mockBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act: Call the controller
	suite.controller.Create(c)

	// Assert: Expect 500 Internal Server Error
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)

	// Verify that the expectation was met
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestGetAllTasks_Success tests fetching tasks successfully
func (suite *TaskControllerTestSuite) TestGetAllTasks_Success() {
	// Arrange: Mock Data
	mockTasks := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task One", Description: "First task"},
		{ID: primitive.NewObjectID(), Title: "Task Two", Description: "Second task"},
	}

	// Expectation: `GetAllTasks` should return the mock tasks
	suite.mockUsecase.On("GetAllTasks", mock.Anything).Return(mockTasks, nil).Once()

	// Create a fake gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/tasks", nil)

	// Act: Call the controller
	suite.controller.GetAllTasks(c)

	// Assert: Check response
	expectedResponse, _ := json.Marshal(gin.H{"tasks": mockTasks})

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), string(expectedResponse), w.Body.String())

	// Verify expectations
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestGetAllTasks_Failure tests fetching tasks when an error occurs
func (suite *TaskControllerTestSuite) TestGetAllTasks_Failure() {
	// Expectation: `GetAllTasks` returns an error
	suite.mockUsecase.On("GetAllTasks", mock.Anything).Return(nil, assert.AnError).Once()

	// Create a fake gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/tasks", nil)

	// Act: Call the controller
	suite.controller.GetAllTasks(c)

	// Assert: Expect 500 Internal Server Error
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)

	// Verify expectations
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestGetTaskByID_Success tests getting a task successfully by ID
func (suite *TaskControllerTestSuite) TestGetTaskByID_Success() {
	// Arrange: Mock Task Data
	mockTask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test Task",
		Description: "Task description",
	}

	// Expectation: `GetTaskByID` should return the mock task
	suite.mockUsecase.On("GetTaskByID", mock.Anything, "task1").Return(&mockTask, nil).Once()

	// Create a fake gin context with a parameter
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "task1"}}
	c.Request = httptest.NewRequest(http.MethodGet, "/tasks/task1", nil)

	// Act: Call the controller
	suite.controller.GetTasksById(c)

	// Assert: Check response
	expectedResponse, _ := json.Marshal(gin.H{"task": mockTask})

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), string(expectedResponse), w.Body.String())

	// Verify expectations
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestGetTaskByID_NotFound tests when a task is not found
func (suite *TaskControllerTestSuite) TestGetTaskByID_NotFound() {
	// Expectation: `GetTaskByID` returns an error (task not found)
	suite.mockUsecase.On("GetTaskByID", mock.Anything, "invalid-id").Return(nil, assert.AnError).Once()

	// Create a fake gin context with a parameter
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "invalid-id"}}
	c.Request = httptest.NewRequest(http.MethodGet, "/tasks/invalid-id", nil)

	// Act: Call the controller
	suite.controller.GetTasksById(c)

	// Assert: Expect 404 Not Found
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)

	// Verify expectations
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestUpdateTask_Success tests successfully updating a task
func (suite *TaskControllerTestSuite) TestUpdateTask_Success() {
	// Arrange: Mock Task Data
	updatedTask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Updated Task",
		Description: "Updated description",
	}

	requestBody, _ := json.Marshal(updatedTask)

	// Expectation: `UpdateTask` should return nil (success)
	suite.mockUsecase.On("UpdateTask", mock.Anything, "task1", &updatedTask).Return(nil).Once()

	// Create a fake gin context with a parameter
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "task1"}}
	c.Request = httptest.NewRequest(http.MethodPut, "/tasks/task1", strings.NewReader(string(requestBody)))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act: Call the controller
	suite.controller.UpdateTask(c)

	// Assert: Check response
	expectedResponse := `{"message":"Task Updated successfully"}`
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())

	// Verify expectations
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestUpdateTask_NotFound tests updating a non-existent task
func (suite *TaskControllerTestSuite) TestUpdateTask_NotFound() {
	// Arrange: Mock Task Data
	updatedTask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Non-existent Task",
		Description: "This task does not exist",
	}

	requestBody, _ := json.Marshal(updatedTask)

	// Expectation: `UpdateTask` returns an error (task not found)
	suite.mockUsecase.On("UpdateTask", mock.Anything, "task2", &updatedTask).Return(assert.AnError).Once()

	// Create a fake gin context with a parameter
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "task2"}}
	c.Request = httptest.NewRequest(http.MethodPut, "/tasks/task2", strings.NewReader(string(requestBody)))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act: Call the controller
	suite.controller.UpdateTask(c)

	// Assert: Expect 404 Not Found
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)

	// Verify expectations
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestDeleteTask_Success tests successfully deleting a task
func (suite *TaskControllerTestSuite) TestDeleteTask_Success() {
	// Arrange: Mock successful deletion
	taskID := primitive.NewObjectID().Hex()
	suite.mockUsecase.On("DeleteTask", mock.Anything, taskID).Return(nil).Once()

	// Create a fake gin context with a task ID parameter
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: taskID}}
	c.Request = httptest.NewRequest(http.MethodDelete, "/tasks/"+taskID, nil)

	// Act: Call the controller
	suite.controller.DeleteTask(c)

	// Assert: Check response
	expectedResponse := `{"message":"Task deleted successfully"}`
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), expectedResponse, w.Body.String())

	// Verify expectations
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestDeleteTask_NotFound tests deleting a non-existent task
func (suite *TaskControllerTestSuite) TestDeleteTask_NotFound() {
	// Arrange: Mock task not found error
	taskID := primitive.NewObjectID().Hex()
	suite.mockUsecase.On("DeleteTask", mock.Anything, taskID).Return(assert.AnError).Once()

	// Create a fake gin context with a task ID parameter
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: taskID}}
	c.Request = httptest.NewRequest(http.MethodDelete, "/tasks/"+taskID, nil)

	// Act: Call the controller
	suite.controller.DeleteTask(c)

	// Assert: Expect 404 Not Found
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)

	// Verify expectations
	suite.mockUsecase.AssertExpectations(suite.T())
}

// Run the test suite
func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}

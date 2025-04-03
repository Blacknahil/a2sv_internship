package usecase_test

import (
	"clean_task_manager_api_tested/domain"
	"clean_task_manager_api_tested/mocks"
	"clean_task_manager_api_tested/usecase"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskUsecaseSuite struct {
	suite.Suite
	repository *mocks.TaskRepositoryInteface
	usecase    domain.TaskUsecaseInteface
}

func (suite *taskUsecaseSuite) SetupTest() {

	repository := new(mocks.TaskRepositoryInteface)
	timeout := time.Second * 10

	// inject the paramaters into the usecase
	usecase := usecase.NewTaskUsecase(repository, timeout)
	suite.repository = repository
	suite.usecase = usecase
}

func TestTaskUsecaseSuite(t *testing.T) {
	suite.Run(t, new(taskUsecaseSuite))
}

func (suite *taskUsecaseSuite) TestCreate_Positive() {

	mockTask := &domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "title 1",
		Description: "this is a task",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	suite.repository.On("Create", mock.Anything, mockTask).Return(mockTask, nil).Once()
	result, err := suite.usecase.Create(context.Background(), mockTask)
	suite.Nil(err)
	suite.Equal(mockTask, result)
}

func (suite *taskUsecaseSuite) TestCreate_Negative() {
	// Arrange
	mockTask := &domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "title 1",
		Description: "this is a task",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	expectedError := errors.New("failed to create task")

	// Mock the repository to return an error
	suite.repository.On("Create", mock.Anything, mockTask).Return(nil, expectedError).Once()

	// Act
	result, err := suite.usecase.Create(context.Background(), mockTask)

	// Assert
	assert.Nil(suite.T(), result)                            // Expect no task to be returned
	assert.EqualError(suite.T(), err, expectedError.Error()) // Expect the error to match
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestGetAllTasks_Positive() {
	// Arrange
	mockTasks := []domain.Task{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 1",
			Description: "Description 1",
			DueDate:     time.Now(),
			Status:      "Pending",
		},
		{
			ID:          primitive.NewObjectID(),
			Title:       "Task 2",
			Description: "Description 2",
			DueDate:     time.Now(),
			Status:      "Completed",
		},
	}

	// Mock the repository to return the tasks
	suite.repository.On("GetAllTasks", mock.Anything).Return(mockTasks, nil).Once()

	// Act
	result, err := suite.usecase.GetAllTasks(context.Background())

	// Assert
	assert.NoError(suite.T(), err)             // Expect no error
	assert.Equal(suite.T(), mockTasks, result) // Expect the returned tasks to match the mock
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestGetAllTasks_Negative() {
	// Arrange
	expectedError := errors.New("failed to retrieve tasks")

	// Mock the repository to return an error
	suite.repository.On("GetAllTasks", mock.Anything).Return(nil, expectedError).Once()

	// Act
	result, err := suite.usecase.GetAllTasks(context.Background())

	// Assert
	assert.Nil(suite.T(), result)                            // Expect no tasks to be returned
	assert.EqualError(suite.T(), err, expectedError.Error()) // Expect the error to match
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestGetTaskByID_Positive() {
	// Arrange
	taskID := primitive.NewObjectID().Hex()
	mockTask := &domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now(),
		Status:      "Pending",
	}

	// Mock the repository to return the task
	suite.repository.On("GetTaskByID", mock.Anything, taskID).Return(mockTask, nil).Once()

	// Act
	result, err := suite.usecase.GetTaskByID(context.Background(), taskID)

	// Assert
	assert.NoError(suite.T(), err)            // Expect no error
	assert.Equal(suite.T(), mockTask, result) // Expect the returned task to match the mock
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestGetTaskByID_Negative() {
	// Arrange
	taskID := primitive.NewObjectID().Hex()
	expectedError := errors.New("task not found")

	// Mock the repository to return an error
	suite.repository.On("GetTaskByID", mock.Anything, taskID).Return(nil, expectedError).Once()

	// Act
	result, err := suite.usecase.GetTaskByID(context.Background(), taskID)

	// Assert
	assert.Nil(suite.T(), result)                            // Expect no task to be returned
	assert.EqualError(suite.T(), err, expectedError.Error()) // Expect the error to match
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestUpdateTask_Positive() {
	// Arrange
	taskID := primitive.NewObjectID().Hex()
	mockTask := &domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     time.Now(),
		Status:      "Completed",
	}

	// Mock the repository to return no error
	suite.repository.On("UpdateTask", mock.Anything, taskID, mockTask).Return(nil).Once()

	// Act
	err := suite.usecase.UpdateTask(context.Background(), taskID, mockTask)

	// Assert
	assert.NoError(suite.T(), err) // Expect no error
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestUpdateTask_Negative() {
	// Arrange
	taskID := primitive.NewObjectID().Hex()
	mockTask := &domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     time.Now(),
		Status:      "Completed",
	}
	expectedError := errors.New("failed to update task")

	// Mock the repository to return an error
	suite.repository.On("UpdateTask", mock.Anything, taskID, mockTask).Return(expectedError).Once()

	// Act
	err := suite.usecase.UpdateTask(context.Background(), taskID, mockTask)

	// Assert
	assert.EqualError(suite.T(), err, expectedError.Error()) // Expect the error to match
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestDeleteTask_Positive() {
	// Arrange
	taskID := primitive.NewObjectID().Hex()

	// Mock the repository to return no error
	suite.repository.On("DeleteTask", mock.Anything, taskID).Return(nil).Once()

	// Act
	err := suite.usecase.DeleteTask(context.Background(), taskID)

	// Assert
	assert.NoError(suite.T(), err) // Expect no error
	suite.repository.AssertExpectations(suite.T())
}

func (suite *taskUsecaseSuite) TestDeleteTask_Negative() {
	// Arrange
	taskID := primitive.NewObjectID().Hex()
	expectedError := errors.New("failed to delete task")

	// Mock the repository to return an error
	suite.repository.On("DeleteTask", mock.Anything, taskID).Return(expectedError).Once()

	// Act
	err := suite.usecase.DeleteTask(context.Background(), taskID)

	// Assert
	assert.EqualError(suite.T(), err, expectedError.Error()) // Expect the error to match
	suite.repository.AssertExpectations(suite.T())
}

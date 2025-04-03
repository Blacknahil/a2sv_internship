package controllers_test

import (
	"bytes"
	"clean_task_manager_api_tested/delivery/controllers"
	"clean_task_manager_api_tested/domain"
	"clean_task_manager_api_tested/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userControllerSuite struct {
	suite.Suite
	mockUserUsecase *mocks.UserUsecaseInterface
	controllers     *controllers.UserController
}

// intialize the test suite

func (suite *userControllerSuite) SetupTest() {
	suite.mockUserUsecase = new(mocks.UserUsecaseInterface)
	suite.controllers = controllers.NewUserController(suite.mockUserUsecase)

}

func (suite *userControllerSuite) TestUserRegister_Postive() {

	mockUser := domain.User{
		Email:    "someone@gmail.com",
		Password: "someone",
	}

	mockBody, _ := json.Marshal(mockUser)

	suite.mockUserUsecase.On("Register", mock.Anything, mockUser).Return(nil).Once()

	// Create a fake gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(mockBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act: Call the controller
	suite.controllers.Register(c)

	// Assert: Check response
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"message":"User registered successfully"}`, w.Body.String())

	// Verify that the expectation was met
	suite.mockUserUsecase.AssertExpectations(suite.T())

}

// TestRegister_InvalidPayload tests user registration with invalid JSON
func (suite *userControllerSuite) TestRegister_InvalidPayload() {
	// Arrange: Invalid JSON (missing required fields)
	invalidBody := `{"undefined":"someone@gmail.com"}`

	// Create a fake gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBufferString(invalidBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act: Call the controller
	suite.controllers.Register(c)

	// Assert: Expect 400 Bad Request
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.JSONEq(suite.T(), `{"error":"Invalid request payload"}`, w.Body.String())

	// Verify that the Register method was NOT called
	suite.mockUserUsecase.AssertNotCalled(suite.T(), "Register", mock.Anything, mock.Anything)
}

// TestRegister_Failure tests user registration when usecase returns an error
func (suite *userControllerSuite) TestRegister_Failure() {
	// Arrange: Mock Data
	mockUser := domain.User{
		Email:    "test@example.com",
		Password: "securepassword",
	}

	mockBody, _ := json.Marshal(mockUser)

	// Expectation: `Register` should return an error
	suite.mockUserUsecase.On("Register", mock.Anything, mockUser).Return(assert.AnError).Once()

	// Create a fake gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(mockBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act: Call the controller
	suite.controllers.Register(c)

	// Assert: Expect 500 Internal Server Error
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
	assert.JSONEq(suite.T(), `{"error":"assert.AnError general error for testing"}`, w.Body.String())

	// Verify that the expectation was met
	suite.mockUserUsecase.AssertExpectations(suite.T())
}

// TestLogin_Success tests login with valid credentials
func (suite *userControllerSuite) TestLogin_Success() {
	// Arrange: Mock Data
	mockLoginRequest := domain.LoginRequest{
		Email:    "test@example.com",
		Password: "securepassword",
	}
	mockLoginResponse := domain.LoginResponse{
		AccessToken: "mocked-jwt-token",
		ID:          primitive.NewObjectID(),
		Role:        "user",
	}

	mockBody, _ := json.Marshal(mockLoginRequest)

	// Expectation: `Login` should be called once and return the mock response
	suite.mockUserUsecase.On("Login", mock.Anything, mockLoginRequest).Return(mockLoginResponse, nil).Once()

	// Create a fake gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(mockBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act: Call the controller
	suite.controllers.Login(c)

	// Assert: Check response
	expectedResponse, _ := json.Marshal(gin.H{
		"message":  "User logged in successfully",
		"Response": mockLoginResponse,
	})

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), string(expectedResponse), w.Body.String())

	// Verify that the expectation was met
	suite.mockUserUsecase.AssertExpectations(suite.T())
}

// TestLogin_InvalidPayload tests login with invalid JSON
func (suite *userControllerSuite) TestLogin_InvalidPayload() {
	// Arrange: Invalid JSON (missing required fields)
	invalidBody := `{"email":"test@example.com"}` // Missing "password"

	// Create a fake gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBufferString(invalidBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act: Call the controller
	suite.controllers.Login(c)

	// Assert: Expect 400 Bad Request
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.JSONEq(suite.T(), `{"error":"Invalid Request Payload"}`, w.Body.String())
}

// TestLogin_Failure tests login with incorrect credentials
func (suite *userControllerSuite) TestLogin_Failure() {
	// Arrange: Mock Data
	mockLoginRequest := domain.LoginRequest{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	mockBody, _ := json.Marshal(mockLoginRequest)

	// Expectation: `Login` should return an error
	suite.mockUserUsecase.On("Login", mock.Anything, mockLoginRequest).Return(domain.LoginResponse{}, assert.AnError).Once()

	// Create a fake gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(mockBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act: Call the controller
	suite.controllers.Login(c)

	// Assert: Expect 400 Bad Request
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	assert.JSONEq(suite.T(), `{"error":"assert.AnError general error for testing"}`, w.Body.String())

	// Verify that the expectation was met
	suite.mockUserUsecase.AssertExpectations(suite.T())
}

func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(userControllerSuite))
}

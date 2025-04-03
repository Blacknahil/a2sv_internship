package router_test

import (
	router "clean_task_manager_api_tested/delivery/routers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mock controllers
type MockUserController struct{ mock.Mock }
type MockTaskController struct{ mock.Mock }

// Implement UserController methods
func (m *MockUserController) Register(c *gin.Context) { m.Called(c) }
func (m *MockUserController) Login(c *gin.Context)    { m.Called(c) }
func (m *MockUserController) Promote(c *gin.Context)  { m.Called(c) }

// Implement TaskController methods
func (m *MockTaskController) Create(c *gin.Context)       { m.Called(c) }
func (m *MockTaskController) GetAllTasks(c *gin.Context)  { m.Called(c) }
func (m *MockTaskController) GetTasksById(c *gin.Context) { m.Called(c) }
func (m *MockTaskController) UpdateTask(c *gin.Context)   { m.Called(c) }
func (m *MockTaskController) DeleteTask(c *gin.Context)   { m.Called(c) }

// Test suite struct
type RouterTestSuite struct {
	suite.Suite
	r                  *gin.Engine
	mockUserController *MockUserController
	mockTaskController *MockTaskController
}

// Setup test suite func (sets up mocks and router)
func (suite *RouterTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.r = gin.Default()
	suite.mockUserController = new(MockUserController)
	suite.mockTaskController = new(MockTaskController)

	router.SetUp(suite.r, suite.mockUserController, suite.mockTaskController)
}

// Run test suite
func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

// Example test for user registration
func (suite *RouterTestSuite) TestUserRegister() {
	suite.mockUserController.On("Register", mock.Anything).Return()

	req := httptest.NewRequest(http.MethodPost, "/users/register", nil)
	w := httptest.NewRecorder()

	suite.r.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockUserController.AssertCalled(suite.T(), "Register", mock.Anything)
}

func (suite *RouterTestSuite) TestUserLogin() {
	// Arrange: Mock the Login method
	suite.mockUserController.On("Login", mock.Anything).Return()

	// Create a fake HTTP request for the login endpoint
	req := httptest.NewRequest(http.MethodPost, "/users/login", nil)
	w := httptest.NewRecorder()

	// Act: Serve the HTTP request
	suite.r.ServeHTTP(w, req)

	// Assert: Check the response status code
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	// Verify that the Login method was called
	suite.mockUserController.AssertCalled(suite.T(), "Login", mock.Anything)
}

func (suite *RouterTestSuite) TestUserPromote_Success() {
	// Arrange: Set up the mock controller to expect the Promote method to be called
	suite.mockUserController.On("Promote", mock.Anything).Return()

	// Simulate a valid user ID for promotion
	userID := primitive.NewObjectID().Hex()

	// Create a fake HTTP request for the promote endpoint
	req := httptest.NewRequest(http.MethodPost, "/users/promote/"+userID, nil)
	w := httptest.NewRecorder()

	// Mock the middleware (authMiddleware and adminMiddleware)
	// We'll allow the request to pass through the middleware
	suite.r.Use(func(c *gin.Context) {
		// Simulate successful auth middleware (no rejection)
		c.Next()
	})

	// Act: Serve the HTTP request (assuming the middlewares allow the request to pass)
	suite.r.ServeHTTP(w, req)

	// Assert: Check the response status code
	assert.Equal(suite.T(), http.StatusOK, w.Code)

	// Verify that the Promote method was called with the correct user ID
	suite.mockUserController.AssertCalled(suite.T(), "Promote", mock.Anything)
}

func (suite *RouterTestSuite) TestUserPromote_Unauthorized() {
	// Simulate a user ID for promotion
	userID := primitive.NewObjectID().Hex()

	// Create a fake HTTP request for the promote endpoint
	req := httptest.NewRequest(http.MethodPost, "/users/promote/"+userID, nil)
	w := httptest.NewRecorder()

	// Assume that middlewares will reject this request, returning a 401 Unauthorized error
	// Act: Serve the HTTP request
	suite.r.ServeHTTP(w, req)

	// Assert: Check that the response code is 401 Unauthorized
	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
}

func (suite *RouterTestSuite) TestUserPromote_NotAdmin() {
	// Simulate a user ID for promotion
	userID := primitive.NewObjectID().Hex()

	// Create a fake HTTP request for the promote endpoint
	req := httptest.NewRequest(http.MethodPost, "/users/promote/"+userID, nil)
	w := httptest.NewRecorder()

	// Assume the user is authenticated but not an admin
	// Act: Serve the HTTP request
	suite.r.ServeHTTP(w, req)

	// Assert: Check that the response code is 403 Forbidden (user not admin)
	assert.Equal(suite.T(), http.StatusForbidden, w.Code)
}

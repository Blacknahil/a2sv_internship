// djjsj
package repository_test

import (
	"clean_task_manager_api_tested/domain"
	"clean_task_manager_api_tested/mocks"
	repository "clean_task_manager_api_tested/repositories"
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
)

// test suite for user repositoryImpl
type userRepositoryImplSuite struct {
	suite.Suite
	repository domain.UserRepositoryInterface
	mockDB     *mocks.DatabaseInterface
}

// SetupTest Intializes the test suit

func (suite *userRepositoryImplSuite) SetupTest() {
	// Create mock database
	suite.mockDB = new(mocks.DatabaseInterface)
}

func (suite *userRepositoryImplSuite) TestUserRegistration_Positive() {
	// Arrange
	mockUser := domain.User{
		Email:    "someone@gmail.com",
		Password: "hashedpassword",
	}

	mockCollection := new(mocks.CollectionInterface)
	suite.mockDB.On("Collection", "users", mock.Anything).Return(mockCollection)

	// Mock the behavior of InsertOne
	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(nil, nil)

	// Mock the behavior of CountDocuments for CheckIfAdmin
	mockCollection.On("CountDocuments", mock.Anything, bson.M{}).Return(int64(0), nil)                // No users in the collection
	mockCollection.On("CountDocuments", mock.Anything, bson.M{"role": "admin"}).Return(int64(0), nil) // No admins

	// Pass the mock database to the repository
	suite.repository = repository.NewUserRepositoryImpl(suite.mockDB, "users")

	// Act
	err := suite.repository.Register(context.Background(), mockUser)

	// Assert
	suite.Nil(err) // Expect no error
	mockCollection.AssertExpectations(suite.T())
}

func TestUserRepositoryImplSuite(t *testing.T) {
	suite.Run(t, new(userRepositoryImplSuite))
}

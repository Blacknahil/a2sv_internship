package usecase_test

import (
	"clean_task_manager_api_tested/domain"
	"clean_task_manager_api_tested/mocks"
	"clean_task_manager_api_tested/usecase"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecaseSuite struct {
	suite.Suite
	repository *mocks.UserRepositoryInterface
	usecase    domain.UserUsecaseInterface
}

func (suite *userUsecaseSuite) SetupTest() {
	// Instantiate the mocked version of the repository
	repository := new(mocks.UserRepositoryInterface)
	timeout := time.Second * 10

	// Inject the repository and timeout into userUsecase
	usecase := usecase.NewUserUsecase(repository, timeout)
	suite.repository = repository
	suite.usecase = usecase
}

func (suite *userUsecaseSuite) TestRegisterUser_Positive() {
	// Define a mock user
	mockUser := domain.User{
		Email:    "someone@gmail.com",
		Password: "someone",
	}

	// Mock the behavior of the repository's Register method
	suite.repository.On("Register", mock.Anything, mockUser).Return(nil).Once()

	// Call the Register method on the usecase
	err := suite.usecase.Register(context.Background(), mockUser)

	// Assert that no error was returned
	suite.Nil(err)

	// Assert that the mock expectations were met
	suite.repository.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestRegisterUser_Negative() {
	// Define a mock user
	mockUser := domain.User{
		Email: "someone@gmail.com", //missing password
	}

	// Mock the behavior of the repository's Register method
	new_error := errors.New("email and password are required")
	suite.repository.On("Register", mock.Anything, mockUser).Return(new_error).Once()

	// Call the Register method on the usecase
	err := suite.usecase.Register(context.Background(), mockUser)

	// Assert that an errror is returned
	suite.EqualError(err, new_error.Error())

	// Assert that the mock expectations were met
	suite.repository.AssertExpectations(suite.T())
}

func (suite *userUsecaseSuite) TestLogin_Positive() {
	mockLoginRequest := domain.LoginRequest{
		Email:    "someone@gmail.com",
		Password: "someone",
	}
	// Convert the string "123" to a primitive.ObjectID
	objectID := primitive.NewObjectID()
	newLoginResponse := domain.LoginResponse{
		AccessToken: "mockAccessToken", // access token
		ID:          objectID,          // id
		Role:        "user",            // role
	}

	// lets mock the behaviour of the repository login method
	suite.repository.On("Login", mock.Anything, mockLoginRequest).Return(newLoginResponse, nil)

	// lets call the login method in the user usecase
	response, err := suite.usecase.Login(context.Background(), mockLoginRequest)

	// assert response and error
	suite.NoError(err)
	suite.Equal(newLoginResponse, response)
}

func (suite *userUsecaseSuite) TestLogin_Negative() {
	mockLoginRequest := domain.LoginRequest{
		Email:    "someone@gmail.com",
		Password: "someone",
	}
	emptyLoginResponse := domain.LoginResponse{}
	newError := errors.New("invalid password or email")

	// lets mock the behaviour of the repository to return an error
	suite.repository.On("Login", mock.Anything, mockLoginRequest).Return(emptyLoginResponse, newError)

	// lets call the login method of the user usecase
	response, err := suite.usecase.Login(context.Background(), mockLoginRequest)

	///assert error
	suite.EqualError(err, newError.Error())
	suite.Equal(emptyLoginResponse, response)

}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(userUsecaseSuite))
}

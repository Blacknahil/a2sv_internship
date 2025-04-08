package repository

import (
	"clean_task_manager_api_tested/domain"
	"clean_task_manager_api_tested/infrastructure"
	"context"
	"errors"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	database   domain.DatabaseInterface
	collection string
}

func (ur *UserRepositoryImpl) Register(c context.Context, user domain.User) error {
	// check if the user does not try to mess up with our database
	collection := ur.database.Collection(ur.collection)

	err := infrastructure.InputValidation(user)
	if err != nil {
		return err
	}

	// hash the password
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return errors.New("internal server error")
	}
	// set default roles and make hashed password as password
	user.Password = string(hashedPassword)

	// should we make the first user admin ?
	isAdmin, err := infrastructure.CheckIfAdmin(c, collection)
	if err != nil {
		return err
	}

	if isAdmin {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	// save to database mongoDB userCollection
	_, err = collection.InsertOne(c, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("email already exists")
		}
		return err
	}

	return nil

}

func (ur *UserRepositoryImpl) Login(c context.Context, loginRequest domain.LoginRequest) (domain.LoginResponse, error) {
	collection := ur.database.Collection(ur.collection)

	err := infrastructure.InputValidation(loginRequest)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	var existingUser domain.User

	// check if the user with the email already exists
	filter := bson.D{{Key: "email", Value: loginRequest.Email}}
	err = collection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err != nil {
		return domain.LoginResponse{}, errors.New("invalid password or email")
	}

	err = infrastructure.ComparePassword(existingUser.Password, loginRequest.Password)
	if err != nil {
		return domain.LoginResponse{}, errors.New("invalid password or email")
	}

	// give the user email and Id to generate the jwt
	expireTimes, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_MINUTES"))
	if err != nil {
		expireTimes = 20 // deafult time for token expiration
	}
	token, err := infrastructure.GenerateToken(existingUser, expireTimes)

	if err != nil {
		return domain.LoginResponse{}, err
	}
	// return the token and/error
	loginResponse := domain.LoginResponse{
		AccessToken: token,
		ID:          existingUser.ID,
		Role:        existingUser.Role,
	}
	return loginResponse, nil

}
func (ur *UserRepositoryImpl) Promote(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.collection)

	// Convert userID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	// Define the filter to find the user by ID
	filter := bson.M{"_id": objectID}

	// Define the update to set the role to "admin"
	update := bson.M{
		"$set": bson.M{
			"role": "admin",
		},
	}

	// Perform the update
	result, err := collection.UpdateOne(c, filter, update)
	if err != nil {
		return err
	}

	// Check if any document was modified
	if result.ModifiedCount == 0 {
		return errors.New("user not found or already an admin")
	}

	return nil
}

func NewUserRepositoryImpl(db domain.DatabaseInterface, collection string) domain.UserRepositoryInterface {

	return &UserRepositoryImpl{
		database:   db,
		collection: collection,
	}
}

// djhfjhsdjhdsjfkdjbjkdsbjksdbj
// dkjakjsdjk
// kjsdkjjkdsjk

package data

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"task_manager_api_auth/constants"
	"task_manager_api_auth/models"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserServices struct {
	UserCollection *mongo.Collection
}

func (us *UserServices) RegisterUser(ctx context.Context, user models.User) error {

	// check if the user does not try to mess up with our database
	err := inputValidation(user)
	if err != nil {
		return err
	}
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("internal server error")
	}
	// set default roles and make hashed password as password
	user.Password = string(hashedPassword)

	isAdmin, _ := us.checkIfAdmin(ctx)
	// could have usefd the err(_) to flag error
	if isAdmin {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	// save to database mongoDB userCollection
	_, err = us.UserCollection.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("email already exists")
		}
		return err
	}

	return nil
}

func (us *UserServices) Login(ctx context.Context, user models.User) (models.LoginResponse, error) {
	err := inputValidation(user)
	if err != nil {
		return models.LoginResponse{}, err
	}

	var existingUser models.User

	// check if the user with the email already exists
	filter := bson.M{"email": user.Email}
	err = us.UserCollection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err != nil {
		return models.LoginResponse{}, errors.New("invalid password or email")
	}

	// compare the hashes
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return models.LoginResponse{}, errors.New("invalid password or email")
	}

	// give the user email and Id to generate the jwt
	expireTimes, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_MINUTES"))
	if err != nil {
		expireTimes = 20 // deafult time for token expiration
	}
	token, err := generateJWT(existingUser, expireTimes)

	if err != nil {
		return models.LoginResponse{}, err
	}
	// return the token and/error
	loginResponse := models.LoginResponse{
		AccessToken: token,
		ID:          existingUser.ID,
		Role:        existingUser.Role,
	}
	return loginResponse, nil
}

func generateJWT(user models.User, expiry int) (string, error) {

	expirationTime := time.Now().Add(time.Duration(expiry) * time.Minute)
	customClaims := models.CustomJWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// get the secret key from the .env file
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET environment variable is not set")
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(constants.SigningMethod), customClaims)
	signedAccessToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}
	return signedAccessToken, nil

}

func inputValidation(user models.User) error {
	// verify that the user only has set the email and password attributes nothing more
	if user.Email == "" || user.Password == "" {
		// other strong and email verifications can be set up here
		return errors.New("email and password are required")
	}

	if !user.ID.IsZero() || user.Role != "" {
		return errors.New("ID and Role must not be set by user")
	}

	return nil

}

func (us *UserServices) checkIfAdmin(ctx context.Context) (bool, error) {
	// check if the user collection is empty
	count, err := us.UserCollection.CountDocuments(ctx, bson.M{})

	if err != nil {
		return false, err
	}
	// if so , return true
	if count == 0 {
		return true, nil
	}
	// else search for anyone who is already an admin
	filter := bson.M{"role": "admin"}

	adminCount, err := us.UserCollection.CountDocuments(ctx, filter)

	if err != nil {
		return false, err
	}

	if adminCount > 0 {
		// return false
		return false, nil
	}

	return true, nil
	// return true to make the first user an admin
}

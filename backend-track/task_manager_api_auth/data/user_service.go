package data

import (
	"context"
	"errors"
	"task_manager_api_auth/models"

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
		return err
	}

	return nil
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

// dsjhjsdjjsd
// / djh

// cjkdfkjf
package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"email" binding:"required"`
}

type LoginResponse struct {
	AccessToken string             `bson:"access_token"`
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Role        string             `bson:"role"`
}

type UserUseCaseInterface interface {
	Register(c context.Context, user User) error
	Login(c context.Context, loginRequest LoginRequest) (LoginResponse, error)
	Promote(c context.Context, userID string) error
}

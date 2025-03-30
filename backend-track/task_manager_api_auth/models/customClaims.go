package models

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomJWTClaims struct {
	UserID primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Email  string             `json:"email"`
	Role   string             `json:"role"`
	jwt.StandardClaims
}

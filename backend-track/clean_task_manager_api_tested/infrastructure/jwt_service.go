package infrastructure

import (
	"clean_task_manager_api_tested/constants"
	"clean_task_manager_api_tested/domain"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(user domain.User, expiry int) (string, error) {

	expirationTime := time.Now().Add(time.Duration(expiry) * time.Minute)
	customClaims := domain.CustomJWTClaims{
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

func ValidateToken(tokenString string) (*domain.CustomJWTClaims, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT secret key not set in the environment variables")
	}

	claims := &domain.CustomJWTClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != constants.SigningMethod {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil

}

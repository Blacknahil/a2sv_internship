package infrastructure

import (
	"clean-task-manager-api/domain"
	"errors"
)

func InputValidation(input interface{}) error {
	switch v := input.(type) {
	case domain.User:
		// Validate User struct
		if v.Email == "" || v.Password == "" {
			return errors.New("email and password are required")
		}
		if !v.ID.IsZero() || v.Role != "" {
			return errors.New("ID and Role must not be set by user")
		}
	case domain.LoginRequest:
		// Validate LoginRequest struct
		if v.Email == "" || v.Password == "" {
			return errors.New("email and password are required")
		}
	default:
		return errors.New("unsupported input type")
	}
	return nil
}

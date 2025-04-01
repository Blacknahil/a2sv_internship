// dfbsdjkbfj
package domain

import "context"

type UserRepositoryInterface interface {
	Register(c context.Context, user User) error
	Login(c context.Context, loginRequest LoginRequest) (LoginResponse, error)
	Promote(c context.Context, userID string) error
}

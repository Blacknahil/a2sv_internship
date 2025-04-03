// dkjskdjjkdsf
package usecase

import (
	"clean_task_manager_api_tested/domain"
	"context"
	"time"
)

type UserUsecase struct {
	userRepositoryInterface domain.UserRepositoryInterface
	contextTimeout          time.Duration
}

func (uu *UserUsecase) Register(c context.Context, user domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepositoryInterface.Register(ctx, user)
}

func (uu *UserUsecase) Login(c context.Context, loginRequest domain.LoginRequest) (domain.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepositoryInterface.Login(ctx, loginRequest)
}

func (uu *UserUsecase) Promote(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepositoryInterface.Promote(ctx, userID)
}

func NewUserUsecase(userRepositoryInterface domain.UserRepositoryInterface, timeout time.Duration) domain.UserUsecaseInterface {

	return &UserUsecase{
		userRepositoryInterface: userRepositoryInterface,
		contextTimeout:          timeout,
	}
}

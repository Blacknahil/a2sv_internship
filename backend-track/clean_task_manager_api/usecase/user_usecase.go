// dkjskdjjkdsf
package usecase

import (
	"clean-task-manager-api/domain"
	"context"
	"time"
)

type userUsecase struct {
	userRepositoryInterface domain.UserRepositoryInterface
	contextTimeout          time.Duration
}

func (uu *userUsecase) Register(c context.Context, user domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepositoryInterface.Register(ctx, user)
}

func (uu *userUsecase) Login(c context.Context, loginRequest domain.LoginRequest) (domain.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepositoryInterface.Login(ctx, loginRequest)
}

func (uu *userUsecase) Promote(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepositoryInterface.Promote(ctx, userID)
}

func NewUserUseCase(userRepositoryInterface domain.UserRepositoryInterface, timeout time.Duration) domain.UserUseCaseInterface {

	return &userUsecase{
		userRepositoryInterface: userRepositoryInterface,
		contextTimeout:          timeout,
	}
}

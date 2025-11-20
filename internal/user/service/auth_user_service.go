package service

import (
	"context"
	"errors"
	password "genpasstore/internal/password/app"
	repository "genpasstore/internal/user/app"
	"genpasstore/internal/user/model"
	"log"
)

type UserService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) Register(ctx context.Context, user model.UserRequest) error {
	if exists, _ := service.repo.ExistsByEmail(ctx, user.Email); exists {
		return errors.New("email already exists")
	}
	hashPassword, err := password.EncodeHashPassword(user.Password)
	if err != nil {
		return err
	}

	err = service.repo.CreateUser(ctx, user.Email, hashPassword)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserService) Login(ctx context.Context, userReq model.UserRequest) (bool, error) {
	if exists, _ := service.repo.ExistsByEmail(ctx, userReq.Email); !exists {
		return false, errors.New("no user with this email was found")
	}
	user, err := service.repo.GetUserByEmail(ctx, userReq.Email)
	if err != nil {
		log.Println("Error while request user from DB")
		return false, err
	}

	check, err := password.VerifyHashPassword(user.Password, userReq.Password)
	if err != nil {
		return false, errors.New("Wrong")
	}

	return check, nil
}

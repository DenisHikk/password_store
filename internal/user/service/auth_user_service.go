package service

import (
	"context"
	"errors"
	password "genpasstore/internal/password/service"
	"genpasstore/internal/user/model"
	"log"
)

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) Register(ctx context.Context, user model.UserRequestRegistry) error {
	if exists, _ := service.repo.ExistsByEmail(ctx, user.Email); exists {
		return errors.New("email already exists")
	}
	hashPassword, err := password.EncodeHashPassword(user.Password)
	if err != nil {
		return err
	}
	hashMasterPassword, err := password.EncodeHashPassword(user.MasterPassword)
	if err != nil {
		return err
	}

	err = service.repo.CreateUser(ctx, user.Email, hashPassword, hashMasterPassword)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserService) Login(ctx context.Context, userReq model.UserRequestsLogin) (bool, error) {
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
		return false, err
	}

	return check, nil
}

package service

import (
	"context"
	"errors"
	password "genpasstore/internal/password/service"
	"genpasstore/internal/user/model"
	userRepo "genpasstore/internal/user/repository"
)

type UserService struct {
	repo userRepo.UserRepository
}

func NewUserService(repo userRepo.UserRepository) *UserService {
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

func (service *UserService) Login(ctx context.Context, user model.UserRequestsLogin) error {
	if exists, _ := service.repo.ExistsByEmail(ctx, user.Email); !exists {
		return errors.New("no user with this email was found")
	}
	return nil
}

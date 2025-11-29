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

func (service *UserService) Register(ctx context.Context, user model.UserRequest) (*model.UserDTO, error) {
	hashPassword, err := password.EncodeHashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	userDTO, err := service.repo.CreateUser(ctx, user.Email, hashPassword)
	if err != nil {
		return nil, err
	}

	return &userDTO, nil
}

func (service *UserService) Login(ctx context.Context, userReq model.UserRequest) (string, error) {
	if exists, _ := service.repo.ExistsByEmail(ctx, userReq.Email); !exists {
		return "", errors.New("no user with this email was found")
	}
	user, err := service.repo.GetUserByEmail(ctx, userReq.Email)
	if err != nil {
		log.Println("Error while request user from DB")
		return "", err
	}

	isSimilar, err := password.VerifyHashPassword(user.Password, userReq.Password)
	if err != nil {
		return "", errors.New("Wrong")
	}

	if !isSimilar {
		return "", errors.New("Wrong")
	}

	return user.ID.String(), nil
}

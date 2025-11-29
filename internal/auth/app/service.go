package auth

import (
	"context"
	"genpasstore/internal/user/model"
)

type AuthService interface {
	Register(ctx context.Context, user model.UserRequest) (*model.UserDTO, error)
	Login(ctx context.Context, user model.UserRequest) (string, error)
}

package auth

import (
	"context"
	"genpasstore/internal/user/model"
)

type AuthService interface {
	Register(ctx context.Context, user model.UserRequest) error
	Login(ctx context.Context, user model.UserRequest) (bool, error)
}

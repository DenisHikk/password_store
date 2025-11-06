package auth

import (
	"context"
	"genpasstore/internal/user/model"
)

type AuthService interface {
	Register(ctx context.Context, user model.UserRequestRegistry) error
	Login(ctx context.Context, user model.UserRequestsLogin) (bool, error)
}

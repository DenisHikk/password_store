package service

import (
	"context"
	"genpasstore/internal/user/model"
)

type Repository interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, email, passwordHash, masterHash string) error
	GetUserByEmail(ctx context.Context, email string) (model.UserDTO, error)
}

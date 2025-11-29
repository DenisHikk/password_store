package user

import (
	"context"
	"database/sql"
	"genpasstore/internal/user/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := repo.db.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`, email).Scan(&exists)
	return exists, err
}

func (repo *UserRepository) CreateUser(ctx context.Context, email, passwordHash string) (model.UserDTO, error) {
	var userDTO model.UserDTO
	err := repo.db.QueryRow(ctx,
		`INSERT INTO users (email, password_hash) VALUES ($1, $2)
		ON CONFLICT (email) DO NOTHING
		RETURNING id, email, created_at`, email, passwordHash).Scan(&userDTO.ID, &userDTO.Email, &userDTO.DateCreate)
	return userDTO, err
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.UserDTO, error) {
	var user model.UserDTO
	rows, err := repo.db.Query(ctx, "SELECT id, email, password_hash, created_at FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.DateCreate)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, sql.ErrNoRows
	}

	return &user, nil
}

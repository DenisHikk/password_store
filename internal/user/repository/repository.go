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

func (repo *UserRepository) CreateUser(ctx context.Context, email, passwordHash string) error {
	_, err := repo.db.Exec(ctx, `INSERT INTO users (email, password_hash) VALUES ($1, $2)`, email, passwordHash)
	return err
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (model.UserDTO, error) {
	var user model.UserDTO
	rows, err := repo.db.Query(ctx, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return model.UserDTO{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return model.UserDTO{}, err
		}
	} else {
		return model.UserDTO{}, sql.ErrNoRows
	}

	return user, nil
}

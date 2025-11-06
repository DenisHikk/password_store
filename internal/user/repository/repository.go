package user

import (
	"context"

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

func (repo *UserRepository) CreateUser(ctx context.Context, email, passwordHash, masterHash string) error {
	_, err := repo.db.Exec(ctx, `INSERT INTO users (email, password_hash, master_hash) VALUES ($1, $2, $3)`, email, passwordHash, masterHash)
	return err
}

package model

import (
	"time"

	"github.com/google/uuid"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	ID         uuid.UUID
	Email      string
	Password   string
	DateCreate time.Time
}

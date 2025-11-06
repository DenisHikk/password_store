package model

import "time"

type UserRequestRegistry struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	MasterPassword string `json:"master_password"`
}

type UserRequestsLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	ID             int32
	Email          string
	Password       string
	MasterPassword string
	DateCreate     time.Time
}

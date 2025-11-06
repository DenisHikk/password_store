package model

type UserRequestRegistry struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	MasterPassword string `json:"master_password"`
}

type UserRequestsLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

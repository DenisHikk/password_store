package auth

import (
	"encoding/json"
	serviceauth "genpasstore/internal/auth/app"
	httpx "genpasstore/internal/httpx/handler"
	"genpasstore/internal/user/model"
	"net/http"
)

type AuthHandler struct {
	service serviceauth.AuthService
}

func NewAuthHandler(service serviceauth.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler *AuthHandler) HandleRegistry(w http.ResponseWriter, r *http.Request) {
	var req model.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid JSON", httpx.ErrorDetails{
			"err": err.Error(),
		})
		return
	}

	err := handler.service.Register(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Error register", httpx.ErrorDetails{
			"err": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req model.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid JSON", httpx.ErrorDetails{
			"err": err.Error(),
		})
		return
	}

	check, err := handler.service.Login(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, http.StatusNotFound, "Error", httpx.ErrorDetails{
			"err": err.Error(),
		})
		return
	}

	if !check {
		httpx.WriteError(w, http.StatusLocked, "Error", httpx.ErrorDetails{
			"err": "Locked account",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
}

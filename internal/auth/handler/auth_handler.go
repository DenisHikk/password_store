package auth

import (
	"encoding/json"
	"errors"
	auth "genpasstore/internal/auth/app"
	httpx "genpasstore/internal/httpx/handler"
	"genpasstore/internal/user/model"
	"net/http"
	"time"
)

type AuthHandler struct {
	service      auth.AuthService
	tokenManager *auth.TokenManager
}

//	{
//	  "access_token": "string",
//	  "token_type": "Bearer",
//	  "expires_in": 0,
//	  "refresh_token": "string"
//	}
type AnswerLogin struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresInRefresh int    `json:"expires_in_refresh"`
}

//	{
//	  "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
//	  "email": "user@example.com",
//	  "created_at": "2025-11-29T20:57:20.927Z"
//	}
type AnswerRegister struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func NewAuthHandler(service auth.AuthService, tokenManager *auth.TokenManager) *AuthHandler {
	return &AuthHandler{service: service, tokenManager: tokenManager}
}

func (handler *AuthHandler) HandleRegistry(w http.ResponseWriter, r *http.Request) {
	var req model.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid JSON", httpx.ErrorDetails{
			"err": err.Error(),
		})
		return
	}

	user, err := handler.service.Register(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Error register", httpx.ErrorDetails{})
		return
	}

	answerRegister := AnswerRegister{
		Id:        user.ID.String(),
		Email:     user.Email,
		CreatedAt: user.DateCreate.String(),
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answerRegister)
}

func (handler *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req model.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid JSON", httpx.ErrorDetails{
			"err": err.Error(),
		})
		return
	}

	if req.Email == "" || req.Password == "" {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid JSON", httpx.ErrorDetails{
			"err": errors.New("access denied"),
		})
		return
	}

	userId, err := handler.service.Login(r.Context(), req)
	if err != nil {
		httpx.WriteError(w, http.StatusNotFound, "Error", httpx.ErrorDetails{
			"err": err.Error(),
		})
		return
	}

	accesToken, err := handler.tokenManager.GenerateAccessToken(userId, 15*time.Minute)
	if err != nil {
		httpx.WriteError(w, http.StatusNotFound, "Error", httpx.ErrorDetails{
			"err": err.Error(),
		})
		return
	}
	refreshToken, err := handler.tokenManager.GenerateRefreshToken(userId, 24*time.Hour)
	if err != nil {
		httpx.WriteError(w, http.StatusNotFound, "Error", httpx.ErrorDetails{
			"err": err.Error(),
		})
		return
	}
	expiresIn := 15 * time.Minute
	expiresInRefresh := 24 * time.Hour
	answerLogin := AnswerLogin{
		AccessToken:      accesToken,
		TokenType:        "Bearer",
		ExpiresIn:        int(expiresIn.Seconds()),
		RefreshToken:     refreshToken,
		ExpiresInRefresh: int(expiresInRefresh.Seconds()),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(answerLogin)
}

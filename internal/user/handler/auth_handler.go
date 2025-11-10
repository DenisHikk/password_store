package auth

import (
	"encoding/json"
	"genpasstore/internal/user/model"
	"log"
	"net/http"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler *AuthHandler) HandleRegistry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method no allowed", http.StatusMethodNotAllowed)
		return
	}
	var req model.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad json", http.StatusBadRequest)
	}

	err := handler.service.Register(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method now allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad json", http.StatusBadRequest)
		return
	}

	check, err := handler.service.Login(r.Context(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if !check {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}

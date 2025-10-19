package handler

import (
	"encoding/json"
	"net/http"
)

type UserRequestRegistry struct {
	Email          string `json:email`
	Password       string `json:password`
	MasterPassword string `json:master_pasword`
}

func HandleRegistry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method no allowd", http.StatusMethodNotAllowed)
		return
	}

	var req UserRequestRegistry
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad json", http.StatusBadRequest)
	}

}

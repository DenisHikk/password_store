package password

import (
	"encoding/json"
	httpx "genpasstore/internal/httpx/handler"
	password "genpasstore/internal/password/service"
	"net/http"
)

type PasswordRequest struct {
	Length    int  `json:"length"`
	UseLower  bool `json:"use_lower"`
	UseUpper  bool `json:"use_upper"`
	UseDigit  bool `json:"use_digit"`
	UseSymbol bool `json:"use_symbol"`
}

type PasswordResponse struct {
	Password string `json:"password"`
}

func HandleGeneratePassword(w http.ResponseWriter, r *http.Request) {
	var req PasswordRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid JSON", httpx.ErrorDetails{
			"err": err.Error(),
		})
		return
	}

	cfg := password.PasswordConfig{
		UseLower:  req.UseLower,
		UseUpper:  req.UseUpper,
		UseDigit:  req.UseDigit,
		UseSymbol: req.UseSymbol,
		Length:    req.Length,
	}

	pw, err := password.GeneratePassword(cfg)
	if err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "Invalid parametrs generate passwird", httpx.ErrorDetails{
			"err": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PasswordResponse{
		Password: pw,
	})
}

package httpx

import (
	"encoding/json"
	"net/http"
)

type ErrorDetails map[string]interface{}

type ErrorResponse struct {
	Error struct {
		Code    int                    `json:"code"`
		Message string                 `json:"message"`
		Details map[string]interface{} `json:"details"`
	} `json:"error"`
}

func WriteError(w http.ResponseWriter, status int, message string, details map[string]interface{}) {
	resp := ErrorResponse{}
	resp.Error.Code = status
	resp.Error.Message = message
	resp.Error.Details = details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

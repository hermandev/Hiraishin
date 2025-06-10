package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hermandev/Hiraishin/apps/node-master/internal/auth"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	payload := auth.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	resp, err := auth.LoginUser(payload.Identity, payload.Password)
	if err != nil {
		http.Error(w, "login failed", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, `{"token":"%s"}`, resp.Token)
}

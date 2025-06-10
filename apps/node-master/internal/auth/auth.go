package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	// Verified        bool   `json:"verified"`
}

type RegisterResponse struct {
	CollectionID    string `json:"collectionId"`
	CollectionName  string `json:"collectionName"`
	ID              string `json:"id"`
	Email           string `json:"email"`
	EmailVisibility bool   `json:"emailVisibility"`
	Verified        bool   `json:"verified"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	Created         string `json:"created"`
	Updated         string `json:"updated"`
}

type LoginResponse struct {
	Token  string          `json:"token"`
	Record json.RawMessage `json:"record"`
}

func LoginUser(email, password string) (*LoginResponse, error) {
	payload := LoginRequest{
		Identity: email,
		Password: password,
	}

	body, _ := json.Marshal(payload)

	resp, err := http.Post("http://127.0.0.1:8090/api/collections/users/auth-with-password", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login failed: %s", resp.Status)
	}

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return nil, fmt.Errorf("invalid response: %w", err)
	}

	return &loginResp, nil
}

func CreateUser(email, password string) (*RegisterResponse, error) {
	payload := RegisterRequest{
		Email:           email,
		Password:        password,
		PasswordConfirm: password,
		// Verified:        true,
	}

	body, _ := json.Marshal(payload)

	resp, err := http.Post("http://127.0.0.1:8090/api/collections/users/records", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("register failed: %s", resp.Status)
	}

	var registerResponse RegisterResponse
	if err := json.NewDecoder(resp.Body).Decode(&registerResponse); err != nil {
		return nil, err
	}

	return &registerResponse, nil
}

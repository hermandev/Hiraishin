package auth_test

import (
	"testing"

	"github.com/hermandev/Hiraishin/apps/node-master/internal/auth"
)

func TestCreateUser_Success(t *testing.T) {
	email := "user@gmail.com"
	password := "user123456"

	resp, err := auth.CreateUser(email, password)
	if err != nil {
		t.Fatalf("Create User failed: %v", err)
	}

	t.Logf("Create User success:  %s", resp.Email)
}

func TestLoginUser_Success(t *testing.T) {
	email := "user@gmail.com"
	password := "user123456"

	resp, err := auth.LoginUser(email, password)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	if resp.Token == "" {
		t.Errorf("Expected token, got empty string")
	}

	t.Logf("Login success: Token = %s", resp.Token)
}

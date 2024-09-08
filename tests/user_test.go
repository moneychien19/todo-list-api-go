package tests

import (
	"testing"

	"github.com/moneychien19/todo-list-api-go/db"
	"github.com/moneychien19/todo-list-api-go/models"
)

func TestCreateUser(t *testing.T) {
	db.InitDB() 
	user := models.User{Name: "John Doe", Email: "john@example.com", Password: "password"}
	token, err := user.CreateUser()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Error("Expected a token, got empty string")
	}
}

func TestLogin(t *testing.T) {
	db.InitDB() 
	user := models.User{Email: "john@example.com", Password: "password"}
	token, err := user.Login()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Error("Expected a token, got empty string")
	}
}
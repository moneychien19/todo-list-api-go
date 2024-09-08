package tests

import (
	"testing"

	"github.com/moneychien19/todo-list-api-go/db"
	"github.com/moneychien19/todo-list-api-go/models"
)

func TestCreateTodo(t *testing.T) {
	db.InitDB() 
	todo := models.Todo{Title: "Test Todo", Description: "Test Description"}
	createdTodo, err := todo.CreateTodo("test@example.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if createdTodo.Title != todo.Title {
		t.Errorf("Expected title %s, got %s", todo.Title, createdTodo.Title)
	}
}

func TestGetTodos(t *testing.T) {
	db.InitDB() 
	todos, err := models.GetTodos(1, 10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(todos) == 0 {
		t.Error("Expected at least one todo")
	}
}

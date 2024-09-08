package models

import (
	"errors"

	"github.com/moneychien19/todo-list-api-go/db"
)

type Todo struct {
	Id int64
	Title string `binding:"required"`
	Description string `binding:"required"`
	CreatedBy string
}

func GetTodos(page, limit int64) ([]Todo, error) {
	query := "SELECT * FROM todos LIMIT ? OFFSET ?"
	rows, err := db.DB.Query(query, limit, (page - 1) * limit)
	if err != nil {
		return nil, err
	}	
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedBy)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (todo *Todo) CreateTodo(userEmail string) (*Todo, error) {
	query := "INSERT INTO todos (title, description, createdBy) VALUES (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(todo.Title, todo.Description, userEmail)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	todo.Id = id
	todo.CreatedBy = userEmail
	
	return todo, nil
}

func (todo *Todo) UpdateTodoById(id int64, userEmail string) (*Todo, error) {
	var createdBy string
	queryCheck := "SELECT createdBy FROM todos WHERE id = ?"
	err := db.DB.QueryRow(queryCheck, id).Scan(&createdBy)
	if err != nil {
		return nil, err
	}
	if createdBy != userEmail {
		return nil, errors.New("unauthorized to update this todo")
	}

	query := "UPDATE todos SET title = ?, description = ? WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Description, id)
	if err != nil {
		return nil, err
	}

	todo.Id = id
	todo.CreatedBy = userEmail
	return todo, nil
}

func DeleteTodoById(id int64, userEmail string) error {
	var createdBy string
	queryCheck := "SELECT createdBy FROM todos WHERE id = ?"
	err := db.DB.QueryRow(queryCheck, id).Scan(&createdBy)
	if err != nil {
		return err
	}
	if createdBy != userEmail {
		return errors.New("unauthorized to delete this todo")
	}
	
	query := "DELETE FROM todos WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
package models

import (
	"github.com/moneychien19/todo-list-api-go/db"
)

type Todo struct {
	Id int64
	Title string `binding:"required"`
	Description string `binding:"required"`
}

func GetTodos() ([]Todo, error) {
	query := "SELECT * FROM todos"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}	
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Description)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (todo *Todo) CreateTodo() (*Todo, error) {
	query := "INSERT INTO todos (title, description) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(todo.Title, todo.Description)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	todo.Id = id
	
	return todo, nil
}

func (todo *Todo) UpdateTodoById(id int64) (*Todo, error) {
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
	return todo, nil
}

func DeleteTodoById(id int64) error {
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
package models

import (
	"github.com/moneychien19/todo-list-api-go/db"
	"github.com/moneychien19/todo-list-api-go/utils"
)

type User struct {
	Name string
	Email string
	Password string
}

func (user *User) CreateUser() (string, error) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(user.Name, user.Email, hashedPassword)
	if err != nil {
		return "", err
	}

	return utils.GenerateToken(user.Email)
}

func (user *User) Login() (string, error) {
	query := "SELECT * FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var dbUser User
	err := row.Scan(&dbUser.Name, &dbUser.Email, &dbUser.Password)
	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(user.Password, dbUser.Password) {
		return "", nil
	}

	return utils.GenerateToken(user.Email)
}
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moneychien19/todo-list-api-go/db"
	"github.com/moneychien19/todo-list-api-go/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")	
}
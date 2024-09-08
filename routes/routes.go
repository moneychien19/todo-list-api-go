package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moneychien19/todo-list-api-go/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authentication)
	authenticated.POST("/todos", createTodos)
	authenticated.PUT("/todos/:id", updateTodos)
	authenticated.DELETE("/todos/:id", deleteTodos)
	authenticated.GET("/todos", getTodos)
	
	server.POST("/register", createUser)
	server.POST("/login", login)
}
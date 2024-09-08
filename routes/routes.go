package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/todos", getTodos)
	server.POST("/todos", createTodos)
	server.PUT("/todos/:id", updateTodos)
	server.DELETE("/todos/:id", deleteTodos)

	server.POST("/register", createUser)
	server.POST("/login", login)
}
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moneychien19/todo-list-api-go/models"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.CreateUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"token": token})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.Login()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": token})
}
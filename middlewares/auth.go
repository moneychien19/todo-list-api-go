package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moneychien19/todo-list-api-go/utils"
)

func Authentication(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	email, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	context.Set("email", email)
	context.Next()
}
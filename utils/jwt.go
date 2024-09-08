package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte("secret"))
}

func ValidateToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	email := claims["email"].(string)
	return email, nil
}
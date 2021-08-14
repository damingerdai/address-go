package utils

import (
	"damingerdai/address/internal/models"
	"fmt"

	jwt "github.com/golang-jwt/jwt"
)

func CreateToken(user *models.User, secretKey []byte, expiresAt int64) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
		"exp":      expiresAt,
	})
	tokenString, err = token.SignedString(secretKey)

	return
}

func VerifyToken(token string, secretKey []byte) bool {
	if len(token) == 0 {
		return false
	}
	t, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})
	return t.Valid
}

package token

import (
	"damingerdai/address/models"

	jwt "github.com/dgrijalva/jwt-go"
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

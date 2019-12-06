package api

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(c *gin.Context) {
	username := c.GetHeader("username")
	password := c.GetHeader("password")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      time.Date(2020, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	message, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, message)
	}

}

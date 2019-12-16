package api

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(c *gin.Context) {
	username := c.GetHeader("username")
	password := c.GetHeader("password")

	if len(username) == 0 {
		c.AbortWithStatusJSON(400, "username is required")
		return
	}

	if len(password) == 0 {
		c.AbortWithStatusJSON(400, "password is required")
		return
	}

	exp := time.Now().Local().Add(time.Hour * time.Duration(1))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      exp.Unix(),
	})

	message, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	} else {
		c.AbortWithStatusJSON(200, message)
	}

}

package api

import (
	"damingerdai/address/internal/models"
	"damingerdai/address/pkg/utils"
	"time"

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

	user := models.User{}
	user.Username = username
	user.Password = password
	secretKey := []byte("damingeridai")
	expiresAt := time.Now().Add(time.Second * 5).Unix()
	token, err := utils.CreateToken(&user, secretKey, expiresAt)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	} else {
		c.AbortWithStatusJSON(200, token)
	}
}

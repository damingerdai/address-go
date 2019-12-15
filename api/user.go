package api

import (
	"damingerdai/address/models"
	service "damingerdai/address/services"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user := models.User{
		Username: "daming",
		Password: "123456",
	}
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)
	if len(user.Username) == 0 || len(user.Password) == 0 {
		c.AbortWithStatusJSON(403, "Bad Request")
	} else {
		service.CreateUser(&user)
		c.JSON(200, user)
	}
}

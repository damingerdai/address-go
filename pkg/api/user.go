package api

import (
	"damingerdai/address/internal/models"
	service "damingerdai/address/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)
	id, err := service.CreateUser(user.Username, user.Password)
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, id)
	}
}

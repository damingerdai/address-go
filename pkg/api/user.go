package api

import (
	"damingerdai/address/internal/models"
	service "damingerdai/address/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

func GetUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(400, errors.Errorf("id '%s' is invalid user id", id))
	} else {
		user, err := service.GetUser(userId)
		if err != nil {
			c.JSON(500, err)
		} else {
			c.JSON(200, user)
		}
	}

}

func ListUsers(c *gin.Context) {
	users, err := service.ListUsers()
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	} else {
		c.AbortWithStatusJSON(200, users)
	}
}

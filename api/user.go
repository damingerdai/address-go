package api

import (
	"damingerdai/address/models"
	service "damingerdai/address/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUser(id)
	if err != nil {
		if err.Error() == "bad request for user id" {
			c.AbortWithStatusJSON(400, "Bad Request")
		} else if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(404, "no user")
		} else {
			c.AbortWithStatusJSON(500, err.Error())
		}
	} else {
		c.AbortWithStatusJSON(200, user)
	}
}

func CreateUser(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)
	if len(user.Username) == 0 || len(user.Password) == 0 {
		c.AbortWithStatusJSON(400, "Bad Request")
	} else {
		service.CreateUser(&user)
		c.JSON(200, user)
	}
}

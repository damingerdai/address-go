package api

import (
	service "damingerdai/address/services"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var hmacSampleSecret = []byte("dam")

func CreateToken(c *gin.Context) {
	username := c.GetHeader("username")
	password := c.GetHeader("password")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	message, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		c.JSON(400, err.Error())
	} else {
		c.JSON(200, message)
	}

}

func ListProvinces(c *gin.Context) {
	c.JSON(200, service.ListProvinces())
}

func GetProvince(c *gin.Context) {
	id := c.Param("id")

	if province, err := service.GetProvince(id); err != nil {
		c.JSON(404, err)
	} else {
		c.JSON(200, province)
	}
}

package api

import (
	service "damingerdai/address/services"
	"os"

	"github.com/gin-gonic/gin"
)

var hmacSampleSecret []byte

func init() {
	secret := os.Getenv("Secret")

	if len(secret) == 0 {
		panic("env Secret is required")
	}

	hmacSampleSecret = []byte(secret)
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

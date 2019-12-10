package api

import (
	service "damingerdai/address/services"

	"github.com/gin-gonic/gin"
)

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

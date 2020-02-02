package api

import (
	service "damingerdai/address/internal/services"
	"github.com/gin-gonic/gin"
)

func ListProvinces(c *gin.Context) {
	if provinces, err := service.ListProvince(); err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, provinces)
	}
}

func GetProvince(c *gin.Context) {
	id := c.Param("id")

	if province, err := service.GetProvince(id); err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, province)
	}

}

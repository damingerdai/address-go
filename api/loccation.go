package api

import (
	service "damingerdai/address/services"

	"github.com/gin-gonic/gin"
)

func GetLocation(c *gin.Context) {
	ip := c.Query("ip")
	if len(ip) > 0 {
		location := service.GetLocationByIP(ip)
		c.JSON(200, location)
	} else {
		c.JSON(401, "params 'ip' is required")
	}
}

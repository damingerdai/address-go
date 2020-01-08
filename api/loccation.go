package api

import "github.com/gin-gonic/gin"
import service "damingerdai/address/services"

func GetLocation(c *gin.Context) {
	ip := c.Query("ip")
	if len(ip) > 0 {
		location := service.GetLocationByIP(ip)
		c.JSON(200, location)
	} else {
		c.JSON(401, "params 'ip' is required")
	}
}

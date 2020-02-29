package api

import (
	service "damingerdai/address/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetIp(c *gin.Context) {
	c.AbortWithStatusJSON(200, c.ClientIP())
}

func GetLocation(c *gin.Context) {
	ip := c.ClientIP()
	if len(ip) == 0 {
		c.AbortWithStatusJSON(500, errors.New("no ip"))
	} else {
		location := service.GetLocationByIP(ip)
		c.AbortWithStatusJSON(200, &location)
	}
}

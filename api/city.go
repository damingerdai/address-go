package api

import "github.com/gin-gonic/gin"
import service "damingerdai/address/services"

func ListCities(c *gin.Context) {
	c.JSON(200, service.ListCities())
}

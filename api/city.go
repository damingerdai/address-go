package api

import "github.com/gin-gonic/gin"
import service "damingerdai/address/services"

func ListCities(c *gin.Context) {
	c.JSON(200, service.ListCities())
}

func GetCity(c *gin.Context) {
	id := c.Param("id")

	if city, err := service.GetCity(id); err != nil {
		c.JSON(404, err)
	} else {
		c.JSON(200, city)
	}
}

package api

import "github.com/gin-gonic/gin"
import service "damingerdai/address/internal/services"

func ListCities(c *gin.Context) {
	if cities, err := service.ListCities(); err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, cities)
	}
}

func GetCity(c *gin.Context) {
	id := c.Param("id")

	if city, err := service.GetCity(id); err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, city)
	}
}

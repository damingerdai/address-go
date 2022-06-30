package api

import (
	"github.com/damingerdai/address-go/internal/config"
	"github.com/damingerdai/address-go/internal/database"
	service "github.com/damingerdai/address-go/internal/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func ListCities(c *gin.Context) {
	db, _ := database.CreateDataSource(config.GetDBConfig())
	trx := db.MustBegin()
	ctx := context.WithValue(c, "trx", trx)
	if cities, err := service.ListCities(ctx); err != nil {
		trx.Rollback()
		c.JSON(500, err)
	} else {
		trx.Commit()
		c.JSON(200, cities)
	}

}

func GetCity(c *gin.Context) {
	db, _ := database.CreateDataSource(config.GetDBConfig())
	trx := db.MustBegin()
	ctx := context.WithValue(c, "trx", trx)
	id := c.Param("id")

	if city, err := service.GetCity(ctx, id); err != nil {
		trx.Rollback()
		c.JSON(500, err)
	} else {
		trx.Commit()
		c.JSON(200, city)
	}
}

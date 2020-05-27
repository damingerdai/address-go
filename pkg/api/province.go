package api

import (
	"context"
	"damingerdai/address/internal/config"
	"damingerdai/address/internal/database"
	service "damingerdai/address/internal/services"

	"github.com/gin-gonic/gin"
)

func ListProvinces(c *gin.Context) {
	db, _ := database.CreateDataSource(config.GetDBConfig())
	trx := db.MustBegin()
	ctx := context.WithValue(c, "trx", trx)
	if provinces, err := service.ListProvince(ctx); err != nil {
		trx.Rollback()
		c.JSON(500, err)
	} else {
		trx.Commit()
		c.JSON(200, provinces)
	}
}

func GetProvince(c *gin.Context) {
	id := c.Param("id")
	db, _ := database.CreateDataSource(config.GetDBConfig())
	trx := db.MustBegin()
	ctx := context.WithValue(c, "trx", trx)
	if province, err := service.GetProvince(ctx, id); err != nil {
		trx.Rollback()
		c.JSON(500, err)
	} else {
		trx.Commit()
		c.JSON(200, province)
	}

}

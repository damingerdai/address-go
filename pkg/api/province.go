package api

import (
	"context"

	"github.com/damingerdai/address-go/internal/config"
	"github.com/damingerdai/address-go/internal/database"
	service "github.com/damingerdai/address-go/internal/services"

	"github.com/gin-gonic/gin"
)

func ListProvinces(c *gin.Context) {
	db, _ := database.CreateDataSource(config.GetDBConfig())
	trx := db.MustBegin()
	ctx := context.WithValue(c, "trx", trx)
	if provinces, err := service.ListProvince(ctx); err != nil {
		trx.Rollback()
		c.JSON(500, err.Error())
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
		c.JSON(500, err.Error())
	} else {
		trx.Commit()
		c.JSON(200, province)
	}

}

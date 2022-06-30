package api

import (
	"context"
	"github.com/damingerdai/address-go/internal/config"
	"github.com/damingerdai/address-go/internal/database"
	"github.com/damingerdai/address-go/internal/models"
	service "github.com/damingerdai/address-go/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func CreateUser(c *gin.Context) {
	user := models.User{}
	c.BindJSON(&user)
	db, _ := database.CreateDataSource(config.GetDBConfig())
	trx := db.MustBegin()
	ctx := context.WithValue(c, "trx", trx)
	id, err := service.CreateUser(ctx, user.Username, user.Password)
	if err != nil {
		trx.Commit()
		c.JSON(500, err)
	} else {
		c.JSON(200, id)
	}
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(400, errors.Errorf("id '%s' is invalid user id", id))
	} else {
		db, _ := database.CreateDataSource(config.GetDBConfig())
		trx := db.MustBegin()
		ctx := context.WithValue(c, "trx", trx)
		user, err := service.GetUser(ctx, userId)
		if err != nil {
			trx.Rollback()
			c.JSON(500, err)
		} else {
			trx.Commit()
			c.JSON(200, user)
		}
	}

}

func ListUsers(c *gin.Context) {
	db, _ := database.CreateDataSource(config.GetDBConfig())
	trx := db.MustBegin()
	ctx := context.WithValue(c, "trx", trx)
	users, err := service.ListUsers(ctx)
	if err != nil {
		trx.Rollback()
		c.AbortWithStatusJSON(500, err)
	} else {
		trx.Commit()
		c.AbortWithStatusJSON(200, users)
	}
}

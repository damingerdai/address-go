package routes

import (
	_ "github.com/damingerdai/address-go/docs"
	"github.com/damingerdai/address-go/pkg/api"
	"github.com/damingerdai/address-go/pkg/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	v1.Use(middleware.ValidateToken())
	{
		v1.GET("ping", api.Ping)
		v1.GET("ip", api.GetIp)
		v1.GET("locatoin", api.GetLocation)
		v1.POST("token", api.CreateToken)
		v1.GET("provinces", api.ListProvinces)
		v1.GET("province/:id", api.GetProvince)
		v1.GET("cities", api.ListCities)
		v1.GET("city/:id", api.GetCity)
		v1.POST("user", api.CreateUser)
		v1.GET("user/:id", api.GetUser)
		v1.GET("users", api.ListUsers)
	}

	return r
}

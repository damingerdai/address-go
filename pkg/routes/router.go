package routes

import (
	"damingerdai/address/pkg/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)
		v1.POST("token", api.CreateToken)
		v1.GET("provinces", api.ListCities)
		v1.GET("province/:id", api.GetProvince)
		v1.GET("cities", api.ListCities)
		v1.GET("city/:id", api.GetCity)
		v1.POST("user", api.CreateUser)
		v1.GET("user/:id", api.GetUser)
	}

	return r
}

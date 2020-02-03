package routes

import (
	"damingerdai/address/api"
	"damingerdai/address/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ValidateToken())
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)
		v1.POST("token", api.CreateToken)
		v1.GET("user/:id", api.GetUser)
		v1.POST("user", api.CreateUser)
		v1.GET("location", api.GetLocation)
	}

	return r
}

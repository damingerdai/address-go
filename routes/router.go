package routes

import (
	"damingerdai/address/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("ping", api.Ping)
	v1.POST("token", api.CreateToken)

	return r
}
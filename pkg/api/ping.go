package api

import "github.com/gin-gonic/gin"

// @Summary ping api
// @Produce  json
// @Success 200 {string} string "成功"
// @Router /api/v1/tags [get]
func Ping(c *gin.Context) {
	c.JSON(200, "pong")
}

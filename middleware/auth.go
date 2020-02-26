package middleware

import (
	"github.com/gin-gonic/gin"

	"damingerdai/address/api"
	utils "damingerdai/address/internal/utils"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() == "/api/v1/token" || c.Request.URL.String() == "/api/v1/ping" || c.Request.URL.String() == "/api/v1/user" {
			c.Next()
		} else {
			tokenStr := c.GetHeader("token")
			if ok := utils.VerifyToken(tokenStr, api.GetHmacSampleSecret()); ok {
				c.Next()
			} else {
				c.AbortWithStatus(403)
				c.JSON(403, "token is invalid")
			}
		}
	}
}

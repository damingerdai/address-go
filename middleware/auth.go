package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"damingerdai/address/api"
	token "damingerdai/address/utils"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() == "/api/v1/token" || c.Request.URL.String() == "/api/v1/ping" {
			c.Next()
		} else {
			tokenStr := c.GetHeader("token")
			fmt.Println(tokenStr)
			if ok := token.VerifyToken(tokenStr, api.GetHmacSampleSecret()); ok {
				c.Next()
			} else {
				c.AbortWithStatus(403)
				c.JSON(403, "token is invalid")
			}
		}
	}
}

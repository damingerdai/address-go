package middleware

import (
	"github.com/gin-gonic/gin"
	"regexp"

	"damingerdai/address/api"
	utils "damingerdai/address/internal/utils"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		if ok, _ := regexp.MatchString("^/api/v1/(token|ping|user)$", url); ok {
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

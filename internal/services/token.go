package service

import (
	"damingerdai/address/internal/config"
	"damingerdai/address/internal/models"
	"damingerdai/address/internal/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateToken(c *gin.Context, user *models.User) {
	secretKey := config.GetHmacSampleSecret()
	expiresAt := time.Now().Add(time.Hour * 1)
	token, err := utils.CreateToken(user, secretKey, expiresAt.Unix())
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	} else {
		response := models.Token{AccessToken: token, ExpiresAt: expiresAt}
		c.AbortWithStatusJSON(200, response)
	}
}

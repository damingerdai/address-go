package service

import (
	"github.com/damingerdai/address-go/internal/config"
	"github.com/damingerdai/address-go/internal/models"
	"github.com/damingerdai/address-go/internal/utils"
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

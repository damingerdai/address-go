package api

import (
	"github.com/gin-gonic/gin"
)

func CreateToken(c *gin.Context) {
	// username := c.GetHeader("username")
	// password := c.GetHeader("password")

	// if len(username) == 0 {
	// 	c.AbortWithStatusJSON(400, "username is required")
	// 	return
	// }

	// if len(password) == 0 {
	// 	c.AbortWithStatusJSON(400, "password is required")
	// 	return
	// }

	// user := models.User{}
	// user.Username = username
	// user.Password = password

	// b, err := service.HasUser(&user)
	// if err != nil {
	// 	c.AbortWithStatusJSON(500, err)
	// 	return
	// }
	// if b {
	// 	exp := time.Now().Local().Add(time.Hour * time.Duration(1))

	// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 		"username": username,
	// 		"password": password,
	// 		"exp":      exp.Unix(),
	// 	})

	// 	message, err := token.SignedString(hmacSampleSecret)

	// 	if err != nil {
	// 		c.AbortWithStatusJSON(500, err.Error())
	// 	} else {
	// 		c.AbortWithStatusJSON(200, message)
	// 	}
	// } else {
	// 	c.AbortWithStatusJSON(404, "no user")
	// }
	c.JSON(200, "Pong")
}

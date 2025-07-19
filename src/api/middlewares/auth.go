package middlewares

import (
	"AuthTemplate/src/models"
	"AuthTemplate/src/resources"
	"AuthTemplate/src/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	jwtKey := strings.Split(tokenString, "Bearer ")
	if len(jwtKey) != 2 {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	jwtData, err := utils.VerifyJWT(jwtKey[1])
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	container := resources.NewContainer(c)

	var user models.User
	if err = container.DB.Preload("Role").Where("id = ?", jwtData).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	c.Set("user", user)

	c.Next()
}

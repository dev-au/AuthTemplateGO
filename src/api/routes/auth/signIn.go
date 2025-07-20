package auth

import (
	"AuthTemplate/src/models"
	"AuthTemplate/src/resources"
	"AuthTemplate/src/utils"
	"github.com/gin-gonic/gin"
)

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func SignIn(c *gin.Context) {
	var req SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(422, gin.H{"error": "Invalid request", "detail": err.Error()})
		return
	}

	container := resources.NewContainer(c)

	var user models.User
	if err := container.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "User does not exist"})
		return
	}

	if !user.IsActive {
		c.JSON(401, gin.H{"error": "User is blocked"})
		return
	}

	jwtKey, _ := utils.GenerateJWT(user.ID)
	c.JSON(200, gin.H{"token": jwtKey})
}

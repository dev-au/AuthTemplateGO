package auth

import (
	"AuthTemplate/src/models"
	"AuthTemplate/src/resources"
	"AuthTemplate/src/utils"
	"github.com/google/uuid"
	"time"

	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     int    `json:"role" binding:"required"`
}

func SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(422, gin.H{"error": "Invalid request", "detail": err.Error()})
		return
	}

	var role models.Role
	container := resources.NewContainer(c)

	if err := container.DB.First(&role, req.Role).Error; err != nil {
		c.JSON(400, gin.H{"error": "Role does not exist"})
		return
	}

	if err := container.DB.Where("email = ?", req.Email).First(&models.User{}).Error; err == nil {
		c.JSON(400, gin.H{"error": "User already exist"})
		return
	}

	verifyKey := uuid.New().String()

	container.Cache.Set("user"+verifyKey, req.Email, time.Minute*5)
	go utils.SendVerifyLink(
		req.Email,
		map[string]interface{}{
			"name":      req.Name,
			"role":      req.Role,
			"password":  req.Password,
			"verifyKey": verifyKey,
		})

	c.JSON(201, gin.H{"success": "Verification link sent"})
}

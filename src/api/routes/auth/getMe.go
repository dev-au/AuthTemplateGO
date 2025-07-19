package auth

import (
	"AuthTemplate/src/models"
	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {

	user := c.MustGet("user").(models.User)

	var role map[string]any = nil
	if user.Role != nil {
		role = map[string]any{
			"id":         user.Role.ID,
			"name":       user.Role.Name,
			"created_at": user.Role.CreatedAt,
			"updated_at": user.Role.UpdatedAt,
		}
	}

	c.JSON(200, gin.H{
		"name":       user.Name,
		"email":      user.Email,
		"is_active":  user.IsActive,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"role":       role,
	})

}

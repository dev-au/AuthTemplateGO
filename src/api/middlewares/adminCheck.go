package middlewares

import (
	"AuthTemplate/src/models"
	"github.com/gin-gonic/gin"
)

func AdminVerify(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	if user.Role != nil {
		c.JSON(403, gin.H{"error": "Permission denied"})
		c.Abort()
		return
	}

	c.Next()
}

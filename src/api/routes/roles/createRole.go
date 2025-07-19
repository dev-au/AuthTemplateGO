package roles

import (
	"AuthTemplate/src/models"
	"AuthTemplate/src/resources"
	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required,min=3,max=255"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	container := resources.NewContainer(c)

	newRole := models.Role{
		Name: req.Name,
	}
	if err := container.DB.Create(&newRole).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, newRole)
}

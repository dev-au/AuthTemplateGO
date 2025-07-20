package roles

import (
	"AuthTemplate/src/models"
	"AuthTemplate/src/resources"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UpdateRole(c *gin.Context) {
	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var req struct {
		Name string `json:"name" binding:"required,min=3,max=255"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	container := resources.NewContainer(c)

	updated := container.DB.Model(&models.Role{}).Where("id = ?", roleID).Update("name", req.Name).RowsAffected

	if updated == 0 {
		c.JSON(400, gin.H{"error": "role not found"})
		return
	}
	c.JSON(200, gin.H{"success": "role updated"})
}

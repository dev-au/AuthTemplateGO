package roles

import (
	"AuthTemplate/src/models"
	"AuthTemplate/src/resources"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DeleteRole(c *gin.Context) {

	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	container := resources.NewContainer(c)

	deletedStmt := container.DB.Delete(&models.Role{}, roleID)
	if deletedStmt.Error != nil {
		c.JSON(404, gin.H{"error": deletedStmt.Error.Error()})
		return
	} else if deletedStmt.RowsAffected == 0 {
		c.JSON(400, gin.H{"error": "role not found"})
		return
	}

	c.JSON(204, gin.H{})
}

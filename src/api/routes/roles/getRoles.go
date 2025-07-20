package roles

import (
	"AuthTemplate/src/models"
	"AuthTemplate/src/resources"
	"github.com/gin-gonic/gin"
	"time"
)

func GetRoles(c *gin.Context) {
	type RoleResponse struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	container := resources.NewContainer(c)

	var roles []RoleResponse
	container.DB.Model(&models.Role{}).Select("id, name").Scan(&roles)

	c.JSON(200, roles)
}

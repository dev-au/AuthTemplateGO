package resources

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Container struct {
	DB    *gorm.DB
	Cache *Cache
}

func NewContainer(c *gin.Context) Container {
	return Container{
		DB: DB.WithContext(c.Request.Context()),
		Cache: &Cache{
			Ctx:    c.Request.Context(),
			Client: RedisClient,
		},
	}
}

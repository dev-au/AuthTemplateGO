package auth

import (
	"AuthTemplate/src/models"
	"AuthTemplate/src/resources"
	"AuthTemplate/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	cipher := c.Param("cipher")
	container := resources.NewContainer(c)
	data, err := utils.DecryptAES(cipher)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad verifier link"})
		return
	}

	fmt.Println(data)

	name := data["name"].(string)
	email := data["email"].(string)
	password := data["password"].(string)
	roleFloat, ok := data["role"].(float64)
	if !ok {
		c.JSON(400, gin.H{"error": "Invalid role format"})
		return
	}
	role := int(roleFloat)

	verifyKey := data["verifyKey"].(string)

	val := container.Cache.Get("user" + email)
	verifyVal, ok := val.(string)
	if !ok || verifyVal != verifyKey {
		c.JSON(400, gin.H{"error": "Invalid or expired verification key"})
		return
	}

	var roleDb models.Role

	if err = container.DB.First(&roleDb, role).Error; err != nil {
		c.JSON(400, gin.H{"error": "Role does not exist"})
		return
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
		IsActive: true,
		RoleID:   &roleDb.ID,
	}

	container.DB.Create(&user)

	jwtKey, _ := utils.GenerateJWT(user.ID)
	c.JSON(200, gin.H{"token": jwtKey})

}

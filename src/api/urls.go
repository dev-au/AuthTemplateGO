package api

import (
	"AuthTemplate/src/api/middlewares"
	"AuthTemplate/src/api/routes/auth"
	"AuthTemplate/src/api/routes/roles"
	"github.com/gin-gonic/gin"
)

func GinEngine() *gin.Engine {
	router := gin.Default()

	authControl := router.Group("/user")
	{
		authControl.POST("/sign-up", auth.SignUp)
		authControl.POST("/sign-in", auth.SignIn)
		authControl.GET("/verify/:cipher", auth.Verify)
		authControl.GET("/get-me", middlewares.AuthMiddleware, auth.GetMe)
	}

	roleControl := router.Group("/role")
	roleControl.Use(middlewares.AuthMiddleware, middlewares.AdminVerify)
	{
		roleControl.POST("", roles.CreateRole)
		roleControl.GET("", roles.GetRoles)
		roleControl.DELETE("/:id", roles.DeleteRole)
		roleControl.PATCH("/:id", roles.UpdateRole)
	}

	return router
}

package routes

import (
	"menu-login/src/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	userController := controllers.NewLoginController()

	v1 := router.Group("/v1")
	{
		v1.GET("/get", userController.FindAll)
		v1.POST("/post", userController.Create)
	}

	return v1
}

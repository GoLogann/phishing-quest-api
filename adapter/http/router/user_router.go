package router

import (
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/")
		userGroup.GET("/:id")
	}
}

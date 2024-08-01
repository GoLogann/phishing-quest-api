package router

import (
	"github.com/gin-gonic/gin"
)

func SetupCategoryRoutes(router *gin.Engine) {
	categoryGroup := router.Group("/categories")
	{
		categoryGroup.POST("/")
		categoryGroup.GET("/")
	}
}

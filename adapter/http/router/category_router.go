package router

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/handler"
)

func SetupCategoryRoutes(router *gin.Engine, categoryHandler *handler.CategoryHandler) {
	categoryGroup := router.Group("api/v1/categories")
	{
		categoryGroup.POST("", categoryHandler.CreateCategory)
		categoryGroup.GET("", categoryHandler.ListCategory)
		categoryGroup.GET("/:id/questions", categoryHandler.ListQuestionsByCategory)
	}
}

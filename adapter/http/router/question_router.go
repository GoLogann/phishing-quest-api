package router

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/handler"
)

func SetupQuestionsRoutes(router *gin.Engine, questionsHandler *handler.QuestionHandler) {
	categoryGroup := router.Group("api/v1/questions")
	{
		categoryGroup.POST("", questionsHandler.CreateQuestion)
		categoryGroup.GET("/:id", questionsHandler.GetQuestion)
		categoryGroup.PUT("/:id", questionsHandler.UpdateQuestion)
		categoryGroup.DELETE("/:id", questionsHandler.DeleteQuestion)
	}
}

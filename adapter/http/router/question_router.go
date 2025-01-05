package router

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/handler"
)

func SetupQuestionRoutes(router *gin.Engine, questionsHandler *handler.QuestionHandler) {
	questionsGroup := router.Group("api/v1/questions")
	{
		questionsGroup.POST("", questionsHandler.CreateQuestion)
		questionsGroup.GET("/:id", questionsHandler.GetQuestion)
		questionsGroup.PUT("/:id", questionsHandler.UpdateQuestion)
		questionsGroup.DELETE("/:id", questionsHandler.DeleteQuestion)
		questionsGroup.GET("/:id/answers", questionsHandler.ListAnswersByQuestion)
	}
}

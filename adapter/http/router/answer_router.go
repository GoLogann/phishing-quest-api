package router

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/handler"
)

func SetupAnswerRoutes(router *gin.Engine, answersHandler *handler.AnswerHandler) {
	answerGroup := router.Group("api/v1/answers")
	{
		answerGroup.POST("", answersHandler.CreateAnswer)
		answerGroup.GET("/:id", answersHandler.GetAnswer)
		answerGroup.PUT("/:id", answersHandler.UpdateAnswer)
		answerGroup.DELETE("/:id", answersHandler.DeleteAnswer)
	}
}

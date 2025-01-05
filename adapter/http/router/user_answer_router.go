package router

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/handler"
)

func SetupUserAnswerRoutes(router *gin.Engine, userAnswerHandler *handler.UserAnswerHandler) {
	userAnswersGroup := router.Group("api/v1/user-answers")
	{
		userAnswersGroup.POST("", userAnswerHandler.CreateUserAnswer)
		userAnswersGroup.GET("", userAnswerHandler.ListUserAnswers)
		userAnswersGroup.GET("/:id", userAnswerHandler.GetUserAnswer)
		userAnswersGroup.PUT("/:id", userAnswerHandler.UpdateUserAnswer)
		userAnswersGroup.DELETE("/:id", userAnswerHandler.DeleteUserAnswer)
	}
}

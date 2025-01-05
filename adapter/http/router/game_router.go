package router

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/handler"
)

func SetupGameRoutes(router *gin.Engine, gameHandler *handler.GameHandler) {
	gameGroup := router.Group("api/v1/game")
	{
		gameGroup.POST("/answer", gameHandler.SubmitAnswer)
	}
}

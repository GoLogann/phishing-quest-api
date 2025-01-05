package router

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/handler"
)

func SetupRankingRoutes(router *gin.Engine, rankingHandler *handler.RankingHandler) {
	gameGroup := router.Group("api/v1/rankings")
	{
		gameGroup.GET("", rankingHandler.GetGlobalRanking)
	}
}

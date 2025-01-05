package http

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/router"
	"phishing-quest/container"
)

func SetupRouter(cont *container.Container) *gin.Engine {
	r := gin.Default()

	router.SetupUserRoutes(r, cont.UserHandler)
	router.SetupCategoryRoutes(r, cont.CategoryHandler)
	router.SetupQuestionRoutes(r, cont.QuestionHandler)
	router.SetupAnswerRoutes(r, cont.AnswerHandler)
	router.SetupUserAnswerRoutes(r, cont.UserAnswerHandler)
	router.SetupGameRoutes(r, cont.GameHandler)
	router.SetupRankingRoutes(r, cont.RankingHandler)
	return r
}

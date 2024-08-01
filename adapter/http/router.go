package http

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/router"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	router.SetupUserRoutes(r)
	router.SetupCategoryRoutes(r)
	//router.SetupQuestionRoutes(r)
	//router.SetupAnswerRoutes(r)

	return r
}

package http

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/router"
	"phishing-quest/container"
)

func SetupRouter(cont *container.Container) *gin.Engine {
	r := gin.Default()

	router.SetupUserRoutes(r, cont.UserHandler)

	return r
}

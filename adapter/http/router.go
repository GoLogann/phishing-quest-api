package http

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/router"
	"phishing-quest/container"
)

func SetupRouter(cont *container.Container) *gin.Engine {
	r := gin.Default()

	// Configuração das rotas com os handlers do container
	router.SetupUserRoutes(r, cont.UserHandler)

	return r
}

package router

import (
	"github.com/gin-gonic/gin"
	"phishing-quest/adapter/http/handler"
)

// SetupUserRoutes configura as rotas relacionadas a usuários
func SetupUserRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	userGroup := router.Group("/api/v1/users")
	{
		userGroup.POST("", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUser)
		// Outras rotas relacionadas a usuários
	}
}

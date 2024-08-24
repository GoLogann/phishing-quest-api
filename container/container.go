package container

import (
	"gorm.io/gorm"
	"phishing-quest/adapter/http/handler"
	"phishing-quest/adapter/repository"
	"phishing-quest/core/usecase"
	"phishing-quest/postgres"
)

type Container struct {
	DB          *gorm.DB
	UserRepo    repository.UserRepository
	UserUseCase *usecase.UserUseCase
	UserHandler *handler.UserHandler
}

func NewContainer() *Container {
	db := postgres.InitDB()
	
	userRepo := repository.NewUserRepository(db)

	userUseCase := usecase.NewUserUseCase(userRepo)

	userHandler := handler.NewUserHandler(userUseCase)

	return &Container{
		DB:          db,
		UserRepo:    userRepo,
		UserUseCase: userUseCase,
		UserHandler: userHandler,
	}
}

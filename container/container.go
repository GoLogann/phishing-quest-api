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
	// Inicializa o banco de dados
	db := postgres.InitDB()

	// Criação dos repositórios
	userRepo := repository.NewUserRepository(db)

	// Criação dos casos de uso
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Criação dos handlers
	userHandler := handler.NewUserHandler(userUseCase)

	return &Container{
		DB:          db,
		UserRepo:    userRepo,
		UserUseCase: userUseCase,
		UserHandler: userHandler,
	}
}

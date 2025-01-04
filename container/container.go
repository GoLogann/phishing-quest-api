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
	UserRepo    *repository.IUserRepository
	UserUseCase *usecase.UserUseCase
	UserHandler *handler.UserHandler

	CategoryRepo    *repository.ICategoryRepository
	CategoryUseCase *usecase.CategoryUseCase
	CategoryHandler *handler.CategoryHandler

	QuestionRepo    *repository.IQuestionRepository
	QuestionUseCase *usecase.QuestionUseCase
	QuestionHandler *handler.QuestionHandler
}

func NewContainer() *Container {
	db := postgres.InitDB()

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	CategoryHandler := handler.NewCategoryHandler(categoryUseCase)

	questionRepo := repository.NewQuestionRepository(db)
	questionUseCase := usecase.NewQuestionUseCase(questionRepo)
	QuestionHandler := handler.NewQuestionHandler(questionUseCase)

	return &Container{
		DB:          db,
		UserRepo:    &userRepo,
		UserUseCase: userUseCase,
		UserHandler: userHandler,

		CategoryRepo:    &categoryRepo,
		CategoryUseCase: categoryUseCase,
		CategoryHandler: CategoryHandler,

		QuestionRepo:    &questionRepo,
		QuestionUseCase: questionUseCase,
		QuestionHandler: QuestionHandler,
	}
}

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

	AnswerRepo    *repository.IAnswerRepository
	AnswerUseCase *usecase.AnswerUseCase
	AnswerHandler *handler.AnswerHandler

	UserScoreRepo *repository.IUserScoreRepository

	UserAnswerRepo    *repository.IUserAnswerRepository
	UserAnswerUseCase *usecase.UserAnswerUseCase
	UserAnswerHandler *handler.UserAnswerHandler

	GameUseCase *usecase.GameUseCase
	GameHandler *handler.GameHandler

	RankingRepo    *repository.IRankingRepository
	RankingUseCase *usecase.RankingUseCase
	RankingHandler *handler.RankingHandler
}

func NewContainer() *Container {
	db := postgres.InitDB()

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	CategoryHandler := handler.NewCategoryHandler(categoryUseCase)

	answerRepo := repository.NewAnswerRepository(db)
	answerUseCase := usecase.NewAnswerUseCase(answerRepo)
	answerHandler := handler.NewAnswerHandler(answerUseCase)

	questionRepo := repository.NewQuestionRepository(db)
	questionUseCase := usecase.NewQuestionUseCase(questionRepo, answerRepo)
	questionHandler := handler.NewQuestionHandler(questionUseCase)

	userScoreRepo := repository.NewUserScoreRepository(db)

	userAnswerRepo := repository.NewUserAnswerRepository(db)
	userAnswerUseCase := usecase.NewUserAnswerUseCase(userAnswerRepo)
	userAnswerHandler := handler.NewUserAnswerHandler(userAnswerUseCase)

	gameUseCase := usecase.NewGameUseCase(answerRepo, userRepo, userScoreRepo)
	gameHandler := handler.NewGameHandler(gameUseCase)

	rankingRepo := repository.NewRankingRepository(db)
	rankingUseCase := usecase.NewRankingUseCase(rankingRepo)
	rankingHandler := handler.NewRankingHandler(rankingUseCase)

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
		QuestionHandler: questionHandler,

		AnswerRepo:    &answerRepo,
		AnswerUseCase: answerUseCase,
		AnswerHandler: answerHandler,

		UserScoreRepo: &userScoreRepo,

		UserAnswerRepo:    &userAnswerRepo,
		UserAnswerUseCase: userAnswerUseCase,
		UserAnswerHandler: userAnswerHandler,

		GameUseCase: gameUseCase,
		GameHandler: gameHandler,

		RankingRepo:    &rankingRepo,
		RankingUseCase: rankingUseCase,
		RankingHandler: rankingHandler,
	}
}

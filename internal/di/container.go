package di

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"phishing-quest/internal/adapters/database"
	handlers2 "phishing-quest/internal/adapters/handlers"
	"phishing-quest/internal/domain/service"
)

type Container struct {
	UserService     *service.UserService
	QuestionService *service.QuestionService
	//AnswerService   *usecase.AnswerService
	//CategoryService *usecase.CategoryService
	UserHandler     *handlers2.UserHandler
	QuestionHandler *handlers2.QuestionHandler
	//AnswerHandler   *handlers.AnswerHandler
	//CategoryHandler *handlers.CategoryHandler
}

func NewContainer(db *sql.DB) *Container {
	// Repositórios

	userRepo := database.NewUserRepository(db)
	questionRepo := database.NewQuestionRepository(db)
	//answerRepo := database.NewAnswerRepository(database)
	//categoryRepo := database.NewCategoryRepository(database)

	// Serviços
	userService := service.NewUserService(userRepo)
	questionService := service.NewQuestionService(questionRepo)
	//answerService := usecase.NewAnswerService(answerRepo)
	//categoryService := usecase.NewCategoryService(categoryRepo)

	// Handlers
	userHandler := handlers2.NewUserHandler(userService)
	questionHandler := handlers2.NewQuestionHandler(questionService)
	//answerHandler := handlers.NewAnswerHandler(answerService)
	//categoryHandler := handlers.NewCategoryHandler(categoryService)

	return &Container{
		UserService:     userService,
		QuestionService: questionService,
		//AnswerService:   answerService,
		//CategoryService: categoryService,
		UserHandler:     userHandler,
		QuestionHandler: questionHandler,
		//AnswerHandler:   answerHandler,
		//CategoryHandler: categoryHandler,
	}
}

func (c *Container) Router() *gin.Engine {
	r := gin.Default()

	// Roteamento de usuários
	r.POST("/users", c.UserHandler.CreateUser)
	r.GET("/users/:id", c.UserHandler.GetUserByID)

	// Roteamento de perguntas
	r.POST("/questions", c.QuestionHandler.CreateQuestion)
	r.GET("/questions/:id", c.QuestionHandler.GetQuestionByID)
	//r.GET("/questions/category/:category_id", c.QuestionHandler.GetQuestionsByCategory)

	// Roteamento de respostas
	//r.POST("/answers", c.AnswerHandler.CreateAnswer)
	//r.GET("/answers/:question_id", c.AnswerHandler.GetAnswersByQuestion)

	// Roteamento de categorias
	//r.POST("/categories", c.CategoryHandler.CreateCategory)
	//r.GET("/categories/:id", c.CategoryHandler.GetCategoryByID)
	//r.GET("/categories", c.CategoryHandler.GetAllCategories)

	return r
}

package di

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"phishing-quest/infra/database"
	"phishing-quest/interfaces/handlers"
	"phishing-quest/usecase"
)

type Container struct {
	UserService     *usecase.UserService
	QuestionService *usecase.QuestionService
	//AnswerService   *usecase.AnswerService
	//CategoryService *usecase.CategoryService
	UserHandler     *handlers.UserHandler
	QuestionHandler *handlers.QuestionHandler
	//AnswerHandler   *handlers.AnswerHandler
	//CategoryHandler *handlers.CategoryHandler
}

func NewContainer(db *sql.DB) *Container {
	// Repositórios

	userRepo := database.NewUserRepository(db)
	questionRepo := database.NewQuestionRepository(db)
	//answerRepo := database.NewAnswerRepository(db)
	//categoryRepo := database.NewCategoryRepository(db)

	// Serviços
	userService := usecase.NewUserService(userRepo)
	questionService := usecase.NewQuestionService(questionRepo)
	//answerService := usecase.NewAnswerService(answerRepo)
	//categoryService := usecase.NewCategoryService(categoryRepo)

	// Handlers
	userHandler := handlers.NewUserHandler(userService)
	questionHandler := handlers.NewQuestionHandler(questionService)
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
	r.GET("/questions/category/:category_id", c.QuestionHandler.GetQuestionsByCategory)

	// Roteamento de respostas
	//r.POST("/answers", c.AnswerHandler.CreateAnswer)
	//r.GET("/answers/:question_id", c.AnswerHandler.GetAnswersByQuestion)

	// Roteamento de categorias
	//r.POST("/categories", c.CategoryHandler.CreateCategory)
	//r.GET("/categories/:id", c.CategoryHandler.GetCategoryByID)
	//r.GET("/categories", c.CategoryHandler.GetAllCategories)

	return r
}

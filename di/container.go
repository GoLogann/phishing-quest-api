package di

type Container struct {
    UserService *usecase.UserService
    UserHandler *handlers.UserHandler
}

func NewContainer(db *sql.DB) *Container {
    userRepo := postgres.NewUserRepository(db)
    userService := usecase.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)

    return &Container{
        UserService: userService,
        UserHandler: userHandler,
    }
}

func (c *Container) Router() *gin.Engine {
    r := gin.Default()

    r.POST("/users", c.UserHandler.CreateUser)
    r.GET("/users/:id", c.UserHandler.GetUserByID)

    return r
}

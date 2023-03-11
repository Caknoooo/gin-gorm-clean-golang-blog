package main

import (
	"os"

	"github.com/Caknoooo/Golang-BLOG/config"
	"github.com/Caknoooo/Golang-BLOG/controller"
	"github.com/Caknoooo/Golang-BLOG/repository"
	"github.com/Caknoooo/Golang-BLOG/routes"
	"github.com/Caknoooo/Golang-BLOG/services"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	var (
		db             *gorm.DB                  = config.SetUpDatabaseConnection()
		jwtService     services.JWTService       = services.NewJWTService()
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    services.UserService      = services.NewUserService(userRepository)
		userController controller.UserController = controller.NewUserController(userService, jwtService)
	)

	server := gin.Default()
	routes.Router(server, userController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	server.Run("127.0.0.1:" + port)
}

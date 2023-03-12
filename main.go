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
)

func main() {
	// if err := godotenv.Load(".env"); err != nil {
	// 	panic(err)
	// }

	var (
		db                *gorm.DB                     = config.SetUpDatabaseConnection()
		jwtService        services.JWTService          = services.NewJWTService()
		userRepository    repository.UserRepository    = repository.NewUserRepository(db)
		userService       services.UserService         = services.NewUserService(userRepository)
		userController    controller.UserController    = controller.NewUserController(userService, jwtService)
		blogRepository    repository.BlogRepository    = repository.NewBlogRepository(db, userRepository)
		blogService       services.BlogService         = services.NewBlogService(blogRepository)
		blogController    controller.BlogController    = controller.NewBlogController(blogService, jwtService)
		commentRepository repository.CommentRepository = repository.NewCommentRepository(db, userRepository)
		commentService    services.CommentService      = services.NewCommentService(commentRepository)
		commentController controller.CommentController = controller.NewCommentController(commentService, jwtService)
	)

	server := gin.Default()
	routes.Router(server, userController, blogController, commentController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}

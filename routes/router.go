package routes

import (
	"github.com/Caknoooo/Golang-BLOG/controller"
	"github.com/Caknoooo/Golang-BLOG/middleware"
	"github.com/Caknoooo/Golang-BLOG/services"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, UserController controller.UserController,  BlogController controller.BlogController, CommentController controller.CommentController, jwtService services.JWTService){
		routes := route.Group("/api/user")
		{
			routes.POST("", UserController.RegisterUser)
			routes.GET("", middleware.Authenticate(jwtService), UserController.GetAllUser)
			routes.POST("/login", UserController.LoginUser)
			routes.DELETE("/", middleware.Authenticate(jwtService), UserController.DeleteUser)
			routes.PUT("/", middleware.Authenticate(jwtService), UserController.UpdateUser)
			routes.GET("/me", middleware.Authenticate(jwtService), UserController.MeUser)
		}

		blogRoutes := route.Group("/api/blog")
		{
			blogRoutes.POST("", middleware.Authenticate(jwtService), BlogController.CreateBlog); // Create
			blogRoutes.GET("", BlogController.GetAllBlog) // Get All Blog
			blogRoutes.GET("/:id", BlogController.GetBlogByID) // Get Blog By ID
			blogRoutes.PUT("/:id", middleware.Authenticate(jwtService), BlogController.UpdateBlog) // Update Blog by ID
			blogRoutes.DELETE("/:id", middleware.Authenticate(jwtService), BlogController.DeleteBlog) // Delete Blog by ID
			blogRoutes.GET("/like/:user_id/:blog_id", middleware.Authenticate(jwtService), BlogController.LikeBlogByBlogID)  
		}

		commentRoutes := route.Group("/api/comment")
		{
			commentRoutes.POST("", middleware.Authenticate(jwtService), CommentController.CreateComment) // Create comment
			commentRoutes.GET("", CommentController.GetAllComment) // Get Comment
			commentRoutes.GET("/:id", CommentController.GetCommentByBlogID) // Get Comment by ID
			commentRoutes.PUT("/:id", middleware.Authenticate(jwtService), CommentController.UpdateComment) // Update Comment
			commentRoutes.DELETE("/:id", middleware.Authenticate(jwtService), CommentController.DeleteComment) // Delete Comment
		}
}
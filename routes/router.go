package routes

import (
	"github.com/Caknoooo/Golang-BLOG/controller"
	"github.com/Caknoooo/Golang-BLOG/middleware"
	"github.com/Caknoooo/Golang-BLOG/services"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, UserController controller.UserController, jwtService services.JWTService){
		routes := route.Group("/api/user")
		{
			routes.POST("", UserController.RegisterUser)
			routes.GET("", middleware.Authenticate(jwtService), UserController.GetAllUser)
			routes.POST("/login", UserController.LoginUser)
			routes.DELETE("/", middleware.Authenticate(jwtService), UserController.DeleteUser)
			routes.PUT("/", middleware.Authenticate(jwtService), UserController.UpdateUser)
			routes.GET("/me", middleware.Authenticate(jwtService), UserController.MeUser)
		}
}
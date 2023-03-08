package main

import (
	"fmt"
	"log"

	"github.com/Caknoooo/Golang-BLOG/config"
	"github.com/Caknoooo/Golang-BLOG/repository"
	"github.com/Caknoooo/Golang-BLOG/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("Hello world")
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	db := config.SetUpDatabaseConnection()

	userRepository := repository.NewUserRepository(db)
	
	server := gin.Default()
	server.Use()

	routes.Router()
}
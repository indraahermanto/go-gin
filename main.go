package main

import (
	"go-gin/controllers"
	"go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	s := gin.Default()
	// s.Use(gin.Logger())
	// s.Use(gin.Recovery())

	// mongo connection
	models.ConnectDB()

	// routes
	s.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
		})
	})

	userController := new(controllers.UserController)
	s.GET("/user", userController.GetUsers)
	s.POST("/user", userController.CreateUser)
	s.GET("/user/:id", userController.DetailUser)
	s.PUT("/user/:id", userController.UpdateUser)
	s.DELETE("/user/:id", userController.DeleteUser)
	s.Run(":5050")
}

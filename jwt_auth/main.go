package main

import (
	"jwtAuth/controllers"
	"jwtAuth/initial"
	"jwtAuth/middleware"
	model "jwtAuth/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	initial.LoadEnvVars()
	initial.ConnectToDb()
	initial.DB.AutoMigrate(&model.User{})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}

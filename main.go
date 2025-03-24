package main

import (
	"biblioteca-api/database"
	"biblioteca-api/models"
	"biblioteca-api/controllers"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	database.InitDB()
	db := database.DB
	db.AutoMigrate(&models.Livro{})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/livros", controllers.GetLivros)
	r.POST("/livros", controllers.CreateLivro)
	r.PUT("/livros/:id", controllers.UpdateLivro)
	r.DELETE("/livros/:id", controllers.DeleteLivro)

	r.Run()
}
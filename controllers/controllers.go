package controllers

import (
	"biblioteca-api/database"
	"biblioteca-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetLivros(c *gin.Context) {
	var livros []models.Livro
	database.DB.Find(&livros)
	c.JSON(http.StatusOK, gin.H{"data": livros})
}

func CreateLivro(c *gin.Context) {
	var input models.Livro
	if err := c.ShouldBindJSON(&input); err != nil { // Se não digitar nada no body, subirá um erro
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	livro := models.Livro{Titulo: input.Titulo, Autor: input.Autor, Ano: input.Ano}
	database.DB.Create(&livro)
	c.JSON(http.StatusCreated, gin.H{"data": livro})
}

func UpdateLivro(c *gin.Context) {
	var livro models.Livro
	if err := database.DB.Where("id = ?", c.Param("id")).First(&livro).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	var input models.Livro
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&livro).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": livro})
}

func DeleteLivro(c *gin.Context) {
	var livro models.Livro
	if err := database.DB.Where("id = ?", c.Param("id")).First(&livro).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}
	
	database.DB.Unscoped().Delete(&livro)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"books-data-back-end/models"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"models"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}
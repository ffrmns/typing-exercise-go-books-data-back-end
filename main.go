package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"books-data-back-end/models"
	"books-data-back-end/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	models.ConnectDatabase()
	r.GET("/books", controllers.FindBooks)
	r.Run()
}
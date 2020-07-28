package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomvc/models"
)

// GetArticle handle request to get an article
func GetArticle(c *gin.Context) {

	articleID := c.Param("articleID")

	thisArticle, err := models.GetArticle(articleID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusFound, gin.H{
			"success": true,
			"message": "Found article successfully",
			"content": thisArticle,
		})
	}
}

// CreateArticle handle request to send create a new article
func CreateArticle(c *gin.Context) {

	title := c.PostForm("title")
	content := c.PostForm("content")

	thisArticle, err := models.CreateArticle(title, content)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Created article successfully",
			"content": thisArticle,
		})
	}
}

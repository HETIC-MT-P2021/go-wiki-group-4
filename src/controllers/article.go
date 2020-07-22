package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomvc/models"
)

// GetArticle handle request to get an article
func (paramHandler *HandleDb) GetArticle(c *gin.Context) {
	dbConnection := paramHandler.DbCon

	articleID := c.Param("articleID")

	thisArticle, err := models.GetArticle(articleID, dbConnection)

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
func (paramHandler *HandleDb) CreateArticle(c *gin.Context) {
	dbConnection := paramHandler.DbCon

	title := c.PostForm("title")
	content := c.PostForm("content")

	thisArticle, err := models.CreateArticle(title, content, dbConnection)

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

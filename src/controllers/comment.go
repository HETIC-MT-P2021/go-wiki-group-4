package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomvc/models"
)

// GetComment handle request to get an article
func GetComment(c *gin.Context) {
	commentID := c.Param("commentID")

	thisComment, err := models.GetComment(commentID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusFound, gin.H{
			"success": true,
			"message": "Found comment successfully",
			"content": thisComment,
		})
	}
}

// CreateComment handle request to send create a new comment
func CreateComment(c *gin.Context) {
	articleID := c.PostForm("articleID")
	content := c.PostForm("content")

	thisComment, err := models.CreateComment(content, articleID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Created comment successfully",
			"content": thisComment,
		})
	}
}

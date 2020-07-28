package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomvc/models"
	"packages.hetic.net/gomvc/views"
)

var view *views.View

// GetArticle handle request to get an article
func GetArticle(c *gin.Context) {
	// articleID := c.Param("articleID")
	// thisArticle, err := models.GetArticle(articleID)
	thisArticle := models.Article{1, "Title 1", "Content 1"}

	view = views.NewView("bootstrap", "views/show_single_article.html")
	view.Render(c.Writer, thisArticle)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"success": false,
	// 		"message": err,
	// 		"content": false,
	// 	})
	// } else {
	// 	view.Render(c.Writer, thisArticle)
	// }
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

// GetLatestArticles fetch latest articles
func GetLatestArticles(c *gin.Context) {
	// latestArticles, err := models.GetArticles()
	latestArticles := []models.Article{
		models.Article{1, "Title 1", "Content 1"},
		models.Article{2, "Title 2", "Content 2"},
		models.Article{3, "Title 3", "Content 3"},
		models.Article{4, "Title 4", "Content 4"},
		models.Article{5, "Title 5", "Content 5"},
	}

	view = views.NewView("bootstrap", "views/show_articles.html")
	view.Render(c.Writer, latestArticles)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"success": false,
	// 		"message": err,
	// 		"content": false,
	// 	})
	// } else {
	// 	view.Render(c.Writer, latestArticles)
	// }
}

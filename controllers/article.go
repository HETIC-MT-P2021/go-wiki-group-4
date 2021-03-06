package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomvc/models"
	utils "packages.hetic.net/gomvc/utils/strategies"
	"packages.hetic.net/gomvc/views"
)

var view *views.View

// GetArticle handle request to get an article
func GetArticle(c *gin.Context) {
	articleID := c.Param("articleID")
	thisArticle, err := models.GetArticle(articleID)

	view = views.NewView("bootstrap", "views/show_single_article.html")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		view.Render(c.Writer, thisArticle)
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

// GetCommentFromArticleID handle request to get an article's comment
func GetCommentFromArticleID(c *gin.Context) {
	articleID := c.Param("articleID")

	comments, err := models.GetCommentFromArticleID(articleID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Found article successfully",
			"content": comments,
		})
	}
}

// GetLatestArticles fetch latest articles
func GetLatestArticles(c *gin.Context) {
	latestArticles, err := models.GetArticles()

	view = views.NewView("bootstrap", "views/show_articles.html")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	} else {
		view.Render(c.Writer, latestArticles)
	}
}

// ExportAs export all artilces in certain format
func ExportAs(c *gin.Context) {
	format := c.DefaultQuery("format", "csv")

	var exportStrategy = utils.ExportStrategyInterface

	if format == "xlsx" {
		exportStrategy = utils.ExportStrategyInterface{
			MIMEType:           utils.XLSXStrategyMIMEType,
			ExportArticlesFile: utils.ExportArticleXLSX,
		}
	} else {
		exportStrategy = utils.ExportStrategyInterface{
			MIMEType:           utils.CSVStrategyMIMEType,
			ExportArticlesFile: utils.ExportArticleCSV,
		}
	}

	strategyContext := utils.ExportStrategyContext{
		SelectedStrategy: exportStrategy,
	}

	response, err := strategyContext.ExportArticles()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err,
			"content": false,
		})
	}

	c.Data(http.StatusOK, fmt.Sprintf("%s; charset=utf-8", response.Type), response.Data)
}

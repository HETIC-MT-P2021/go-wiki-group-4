package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gomvc/controllers"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API is running successfully",
		"success": true,
	})
}

// StartRouter will launch the web server
func StartRouter(apiPort string) {
	router := gin.New()

	public := router.Group("/")
	{
		public.GET("/", controllers.GetLatestArticles)

		// Articles
		public.GET("/articles/:articleID", controllers.GetArticle)
		public.POST("/articles", controllers.CreateArticle)
		public.GET("/export", controllers.ExportAs)

		// Comments
		public.POST("/comment", controllers.CreateComment)
	}

	router.Run(":" + apiPort)
}

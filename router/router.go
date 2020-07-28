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
		public.GET("/", healthCheck)
	}

	apiRoutes := router.Group("/api")
	{
		//Articles

		apiRoutes.GET("/article/:articleID", controllers.GetArticle)

		apiRoutes.POST("/article", controllers.CreateArticle)

		// Comments

		apiRoutes.GET("/comment/:commentID", controllers.GetComment)

		apiRoutes.POST("/comment", controllers.CreateComment)
	}

	router.Run(":" + apiPort)
}

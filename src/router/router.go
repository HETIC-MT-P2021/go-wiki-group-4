package router

import (
	"database/sql"
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
func StartRouter(apiPort string, dbCon *sql.DB) {
	router := gin.New()

	DbHandler := new(controllers.HandleDb)
	DbHandler.DbCon = dbCon

	public := router.Group("/")
	{
		public.GET("/", healthCheck)

	}

	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/", healthCheck)
	}

	router.Run(":" + apiPort)
}

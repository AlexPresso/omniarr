package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"omniarr/internal/api/handlers"
	"strings"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/health", handlers.HealthHandler)

		medias := api.Group("/medias")
		{
			medias.GET("/search", handlers.MediaSearchHandler)
			medias.GET("/:media", handlers.MediaDetailsHandler)
		}

		downloads := api.Group("/downloads")
		{
			downloads.POST("/query", handlers.DownloadsSearchHandler)
			downloads.POST("/queue", handlers.QueueDownloadHandler)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		if c.Request.Method == "GET" && !strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.File("./web/dist/index.html")
		} else {
			c.String(http.StatusNotFound, "404 not found")
		}
	})
}

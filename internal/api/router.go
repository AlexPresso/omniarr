package api

import (
	"github.com/gofiber/fiber/v2"
	"omniarr/internal/api/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/", "./web/build")

	api := app.Group("/api")
	api.Get("/health", handlers.HealthHandler)

	medias := api.Group("/medias")
	medias.Get("/search", handlers.MediaSearchHandler)
	medias.Get("/:media", handlers.MediaDetailsHandler)

	downloads := api.Group("/downloads")
	downloads.Get("/search", handlers.DownloadsSearchHandler)
	downloads.Post("/queue", handlers.QueueDownloadHandler)
}

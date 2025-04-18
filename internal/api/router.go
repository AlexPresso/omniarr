package api

import (
	"github.com/gofiber/fiber/v2"
	"omniarr/internal/api/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/", "./web/dist")

	api := app.Group("/api")
	api.Get("/health", handlers.HealthHandler)

	medias := api.Group("/medias")
	medias.Get("/search", handlers.MediaSearchHandler)
	medias.Get("/:media", handlers.MediaDetailsHandler)

	downloads := api.Group("/downloads")
	downloads.Post("/query", handlers.DownloadsSearchHandler)
	downloads.Post("/queue", handlers.QueueDownloadHandler)
}

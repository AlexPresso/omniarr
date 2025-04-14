package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"omniarr/internal/api/response"
	"omniarr/internal/core/download"
	"omniarr/internal/core/media"
)

func DownloadsSearchHandler(c *fiber.Ctx) error {
	ctx := context.Background()

	query := c.Query("q")
	if query == "" {
		return response.Fail(c, "Missing ?q= param", fiber.StatusBadRequest)
	}

	mediaType := c.Query("type")
	if mediaType == "" {
		return response.Fail(c, "Missing ?type= param", fiber.StatusBadRequest)
	}

	results, err := download.Search(ctx, query, media.Type(mediaType))
	if err != nil {
		return response.Fail(c, "")
	}

	return response.OK(c, results)
}

func QueueDownloadHandler(c *fiber.Ctx) error {
	ctx := context.Background()

	var req download.QueueDownloadRequest
	if err := c.BodyParser(&req); err != nil {
		log.Error(err)
		return response.Fail(c, "Failed to parse request", fiber.StatusBadRequest)
	}

	if err := download.QueueDownload(ctx, req.Url); err != nil {
		log.Error(err)
		return response.Fail(c, "Error while queuing download")
	}

	return response.OK(c, nil)
}

package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"omniarr/internal/api/response"
	"omniarr/internal/core/torrent"
)

func DownloadsSearchHandler(c *fiber.Ctx) error {
	ctx := context.Background()

	query := c.Query("q")
	if query == "" {
		return response.Fail(c, "Missing ?q= param", fiber.StatusBadRequest)
	}

	results, err := torrent.Search(ctx, query)
	if err != nil {
		return response.Fail(c, "")
	}

	return response.OK(c, results)
}

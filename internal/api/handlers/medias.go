package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"omniarr/internal/api/response"
	"omniarr/internal/core/media"
	"strconv"
	"strings"
)

func MediaSearchHandler(c *fiber.Ctx) error {
	ctx := context.Background()

	mediaTypes := []media.Type{"movie", "tv"}
	mediaType := c.Query("type")
	if mediaType != "" {
		mediaTypes = []media.Type{"movie"}
	}

	query := c.Query("q")
	if query == "" {
		return response.Fail(c, "Missing ?q= param", fiber.StatusBadRequest)
	}

	results, err := media.Search(ctx, query, mediaTypes)
	if err != nil {
		log.Error("Error searching medias: %v", err)
		return response.Fail(c, "")
	}

	return response.OK(c, results)
}

func MediaDetailsHandler(c *fiber.Ctx) error {
	ctx := context.Background()
	mediaSplit := strings.Split(c.Params("media"), ":")
	if len(mediaSplit) != 2 {
		return response.Fail(c, "Media ID should be <type>:<id>")
	}

	mediaType := media.Type(mediaSplit[0])
	idString := mediaSplit[1]

	id, err := strconv.Atoi(idString)
	details, err := media.GetDetails(ctx, id, mediaType)
	if err != nil {
		log.Error("Error getting media details: %v", err)
		return response.Fail(c, "")
	}

	return response.OK(c, details)
}

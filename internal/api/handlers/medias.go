package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"omniarr/internal/api/response"
	"omniarr/internal/core/media"
	"strconv"
)

func MediaSearchHandler(c *fiber.Ctx) error {
	ctx := context.Background()

	query := c.Query("q")
	if query == "" {
		return response.Fail(c, "Missing ?q= param", fiber.StatusBadRequest)
	}

	results, err := media.Search(ctx, query)
	if err != nil {
		log.Printf("Error searching medias: %v", err)
		return response.Fail(c, "")
	}

	return response.OK(c, results)
}

func MediaDetailsHandler(c *fiber.Ctx) error {
	ctx := context.Background()
	idString := c.Params("id")

	id, err := strconv.Atoi(idString) //This will change when Omniarr supports different providers
	details, err := media.GetDetails(ctx, id)
	if err != nil {
		log.Printf("Error getting media details: %v", err)
		return response.Fail(c, "")
	}

	return response.OK(c, details)
}

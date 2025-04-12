package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"omniarr/internal/api"
	"omniarr/internal/config"
)

func main() {
	app := fiber.New()

	api.SetupRoutes(app)

	addr := fmt.Sprintf(":%s", config.AppConfig.Port)
	log.Printf("🚀 Omniarr is running at http://localhost%s\n", addr)

	if err := app.Listen(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

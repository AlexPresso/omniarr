package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"omniarr/internal/api"
	"omniarr/internal/config"
)

func main() {
	router := gin.Default()

	api.SetupRoutes(router)

	addr := fmt.Sprintf(":%s", config.AppConfig.Port)
	log.Printf("🚀 Omniarr is running at http://localhost%s\n", addr)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

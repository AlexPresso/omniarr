package handlers

import (
	"github.com/gin-gonic/gin"
	"omniarr/internal/api/response"
)

func HealthHandler(c *gin.Context) {
	response.OK(c, map[string]interface{}{
		"status":  "ok",
		"version": "1.0.0",
	})
}

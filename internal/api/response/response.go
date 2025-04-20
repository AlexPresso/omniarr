package response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error"`
}

type StreamResult[T any] struct {
	Data  T
	Error error
}

func (r Response) Error() string {
	return r.ErrorMessage
}

func JSON(c *gin.Context, status int, data interface{}, errMsg string) {
	c.JSON(status, Response{
		Data:         data,
		ErrorMessage: errMsg,
	})
}

func OK(c *gin.Context, data interface{}) {
	JSON(c, http.StatusOK, data, "")
}

func Fail(c *gin.Context, msg string, status ...int) {
	code := http.StatusInternalServerError
	if len(status) > 0 {
		code = status[0]
	}

	message := "An internal server error occurred (check logs)."
	if len(msg) > 0 {
		message = msg
	}

	JSON(c, code, nil, message)
}

func Stream[T any](c *gin.Context, stream <-chan StreamResult[T]) {
	c.Writer.Header().Set("Content-Type", "application/x-ndjson")
	c.Writer.Header().Set("Content-Encoding", "identity")
	c.Status(http.StatusOK)

	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.String(http.StatusInternalServerError, "Streaming not supported")
		return
	}

	enc := json.NewEncoder(c.Writer)
	for result := range stream {
		if result.Error != nil {
			continue
		}
		if err := enc.Encode(result.Data); err != nil {
			continue
		}
		flusher.Flush()
	}
}

package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"omniarr/internal/api/response"
	"omniarr/internal/core/media"
	"strings"
	"sync"
)

func MediaSearchHandler(c *gin.Context) {
	ctx := context.Background()

	mediaTypes := []media.Type{"movie", "tv", "book"}
	mediaType := c.Query("type")
	if mediaType != "" {
		mediaTypes = []media.Type{media.Type(mediaType)}
	}

	query := c.Query("q")
	if query == "" {
		response.Fail(c, "Missing ?q= param", http.StatusBadRequest)
		return
	}

	resultChan := make(chan response.StreamResult[*media.Media], len(mediaTypes))
	go func() {
		var lock sync.WaitGroup

		for _, t := range mediaTypes {
			t := t
			lock.Add(1)

			go func(mt media.Type) {
				defer lock.Done()

				results, err := media.Search(ctx, query, mt)
				if err != nil {
					resultChan <- response.StreamResult[*media.Media]{Error: err}
					return
				}

				for _, result := range results {
					resultChan <- response.StreamResult[*media.Media]{Data: result}
				}
			}(t)
		}

		lock.Wait()
		close(resultChan)
	}()

	response.Stream(c, resultChan)
}

func MediaDetailsHandler(c *gin.Context) {
	ctx := context.Background()
	mediaSplit := strings.Split(c.Param("media"), ":")
	if len(mediaSplit) != 2 {
		response.Fail(c, "Media ID should be <type>:<id>")
		return
	}

	mediaType := media.Type(mediaSplit[0])
	idString := mediaSplit[1]

	details, err := media.GetDetails(ctx, idString, mediaType)
	if err != nil {
		response.Fail(c, "")
	}

	response.OK(c, details)
}

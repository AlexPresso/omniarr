package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"omniarr/internal/api/response"
	"omniarr/internal/core/download"
	"omniarr/internal/core/media"
	"sync"
)

func DownloadsSearchHandler(c *gin.Context) {
	ctx := context.Background()
	var search download.SearchQuery

	if err := c.ShouldBindJSON(&search); err != nil {
		response.Fail(c, "Failed to parse request", http.StatusBadRequest)
		return
	}

	if search.Type == "" {
		response.Fail(c, "Missing ?type= param", http.StatusBadRequest)
		return
	}
	if search.Title == "" {
		response.Fail(c, "Missing ?title= param", http.StatusBadRequest)
		return
	}

	queries := download.GenerateQueries(search)
	resultChan := make(chan response.StreamResult[download.Download], len(queries))
	go func() {
		var lock sync.WaitGroup

		for _, q := range queries {
			q := q
			lock.Add(1)

			go func(query string) {
				defer lock.Done()

				results, err := download.Search(ctx, media.Type(search.Type), q)
				if err != nil {
					resultChan <- response.StreamResult[download.Download]{Error: err}
					return
				}

				for _, result := range results {
					resultChan <- response.StreamResult[download.Download]{Data: result}
				}
			}(q)
		}

		lock.Wait()
		close(resultChan)
	}()

	response.Stream(c, resultChan)
}

func QueueDownloadHandler(c *gin.Context) {
	ctx := context.Background()

	var req download.QueueDownloadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, "Failed to parse request", http.StatusBadRequest)
		return
	}

	if err := download.QueueDownload(ctx, req.Url); err != nil {
		response.Fail(c, "Error while queuing download")
		return
	}

	response.OK(c, nil)
}

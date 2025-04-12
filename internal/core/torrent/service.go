package torrent

import (
	"context"
	"github.com/webtor-io/go-jackett"
	"omniarr/internal/config"
)

func Search(ctx context.Context, query string, categories []uint) ([]Result, error) {
	resp, err := config.JackettClient.Fetch(ctx, &jackett.FetchRequest{
		Query:      query,
		Categories: categories,
	})

	if err != nil {
		return nil, err
	}

	return mapToSearchResults(resp.Results), nil
}

func mapToSearchResults(results []jackett.Result) []Result {
	var sanitized []Result

	for _, r := range results {
		sanitized = append(sanitized, Result{
			Guid:        r.Guid,
			Title:       r.Title,
			Size:        r.Size,
			Seeders:     r.Seeders,
			Leechers:    r.Peers - r.Seeders,
			Category:    r.CategoryDesc,
			Indexer:     r.Tracker,
			PublishDate: r.PublishDate.String(),
		})
	}

	return sanitized
}

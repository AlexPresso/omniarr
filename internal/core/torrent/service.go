package torrent

import (
	"context"
	"github.com/webtor-io/go-jackett"
	"omniarr/internal/client"
	"regexp"
	"strings"
)

func Search(ctx context.Context, query string) ([]Torrent, error) {
	req := &jackett.FetchRequest{
		Query: query,
	}

	resp, err := client.JackettClient.Fetch(ctx, req)
	if err != nil {
		return nil, err
	}

	return mapToSearchResults(resp.Results), nil
}

func NormalizeQuery(title, year string) string {
	query := title + " " + year

	re := regexp.MustCompile(`[,:;"\-]`)
	query = re.ReplaceAllString(query, "")

	query = strings.Join(strings.Fields(query), " ")
	return query
}

func mapToSearchResults(results []jackett.Result) []Torrent {
	torrentsMap := make(map[string]Torrent)
	for _, r := range results {
		torrentsMap[r.Guid] = Torrent{
			Guid:        r.Guid,
			Title:       r.Title,
			Size:        r.Size,
			Seeders:     r.Seeders,
			Leechers:    r.Peers,
			Category:    r.CategoryDesc,
			Indexer:     r.Tracker,
			PublishDate: r.PublishDate.String(),
		}
	}

	torrents := make([]Torrent, 0, len(torrentsMap))
	for _, t := range torrentsMap {
		torrents = append(torrents, t)
	}

	return torrents
}

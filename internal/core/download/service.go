package download

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"github.com/hekmon/transmissionrpc"
	"github.com/webtor-io/go-jackett"
	"omniarr/internal/client"
	"omniarr/internal/config"
	"omniarr/internal/core/media"
	"strings"
)

var categories = map[media.Type][]uint{
	"movie": {
		2000, // Movies (general)
		2010, // DVDR
		2020, // HD
		2030, // UHD / 4K
		2040, // BluRay
		2045, // Remux
		2050, // WEB-DL
		2060, // x264
		2070, // x265 / HEVC
		2080, // 3D
		2090, // Other
	},
	"tv": {
		5000, // TV (general)
		5010, // TV - SD
		5020, // TV - HD
		5030, // TV - UHD / 4K
		5040, // TV - BluRay
		5045, // TV - Remux
		5050, // TV - WEB-DL
		5060, // TV - x264
		5070, // TV - x265 / HEVC
		5080, // TV - Other
	},
}

func Search(ctx context.Context, mediaType media.Type, query string) ([]Download, error) {
	queryCategories := append(categories[mediaType], 8000)

	log.Info("Searching for %s", query)
	req := &jackett.FetchRequest{
		Query:      query,
		Categories: queryCategories,
	}

	resp, err := client.Jackett.Fetch(ctx, req)
	if err != nil {
		log.Error("Error fetching results from Jackett: %v", err)
		return nil, err
	}

	return filterAndMapToSearchResults(resp.Results), nil
}

func QueueDownload(ctx context.Context, url string) error {
	url = strings.ReplaceAll(url, "<jackettApiKey>", config.AppConfig.JackettApiKey)
	_, err := client.Transmission.TorrentAdd(&transmissionrpc.TorrentAddPayload{
		Filename: &url,
	})
	if err != nil {
		return err
	}

	return nil
}

func GenerateQueries(query SearchQuery) []string {
	queries := media.MakeAlternateTitles(query.Title)
	queries = append(queries, media.MakeAlternateTitles(query.OriginalTitle)...)

	return queries
}

func filterAndMapToSearchResults(results []jackett.Result) []Download {
	torrentsMap := make(map[string]Download)
	for _, r := range results {
		if r.Seeders == 0 {
			continue
		}

		link := strings.ReplaceAll(r.Link, config.AppConfig.JackettApiKey, "<jackettApiKey>")
		magnetUri := strings.ReplaceAll(r.MagnetUri, config.AppConfig.JackettApiKey, "<jackettApiKey>")

		torrentsMap[r.Guid] = Download{
			Title:       r.Title,
			Size:        r.Size,
			Seeders:     r.Seeders,
			Leechers:    r.Peers,
			Indexer:     r.Tracker,
			PublishDate: r.PublishDate.String(),
			Link:        link,
			MagnetUri:   magnetUri,
		}
	}

	torrents := make([]Download, 0, len(torrentsMap))
	for _, t := range torrentsMap {
		torrents = append(torrents, t)
	}

	return torrents
}

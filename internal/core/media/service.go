package media

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
)

var mediaFetchers = map[Type]Fetcher{
	"movie": MovieFetcher{},
	"tv":    TVFetcher{},
	"book":  BookFetcher{},
}

func Search(ctx context.Context, query string, mediaTypes []Type) ([]*Media, error) {
	var medias []*Media

	for _, mediaType := range mediaTypes {
		fetcher, ok := mediaFetchers[mediaType]
		if !ok {
			return nil, fmt.Errorf("unsupported media type: %s", mediaType)
		}

		media, err := fetcher.Search(ctx, query)
		if err != nil {
			log.Error(err)
			continue
		}

		medias = append(medias, media...)
	}

	return medias, nil
}

func GetDetails(ctx context.Context, id string, mediaType Type) (*Media, error) {
	fetcher, ok := mediaFetchers[mediaType]
	if !ok {
		return nil, fmt.Errorf("unsupported media type: %s", mediaType)
	}

	media, err := fetcher.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}

	return media, nil
}

package media

import (
	"context"
	"fmt"
)

var mediaFetchers = map[Type]Fetcher{
	"movie": MovieFetcher{},
	"tv":    TVFetcher{},
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
			return nil, err
		}

		medias = append(medias, media...)
	}

	return medias, nil
}

func GetDetails(ctx context.Context, id int, mediaType Type) (*Media, error) {
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

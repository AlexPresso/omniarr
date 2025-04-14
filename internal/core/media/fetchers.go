package media

import (
	"context"
	"github.com/jinzhu/copier"
	"omniarr/internal/client"
)

type Fetcher interface {
	Fetch(ctx context.Context, id int) (*Media, error)
	Search(ctx context.Context, query string) ([]*Media, error)
}

type MovieFetcher struct{}
type TVFetcher struct{}
type MusicFetcher struct{}
type BookFetcher struct{}

func (MovieFetcher) Fetch(ctx context.Context, id int) (*Media, error) {
	movie, err := client.TMDB.GetMovieDetails(id, client.TMDBDefaultOptions)
	if err != nil {
		return nil, err
	}

	return ToMedia(&MovieWrapper{movie}), nil
}

func (MovieFetcher) Search(ctx context.Context, query string) ([]*Media, error) {
	movies, err := client.TMDB.GetSearchMovies(query, client.TMDBDefaultOptions)
	if err != nil {
		return nil, err
	}

	var medias []*Media
	for _, movie := range movies.Results {
		var media TMDBMediaLight
		err := copier.Copy(&media, movie)
		if err != nil {
			return nil, err
		}

		media.Type = "movie"
		medias = append(medias, ToMedia(&media))
	}

	return medias, nil
}

func (TVFetcher) Fetch(ctx context.Context, id int) (*Media, error) {
	tvShow, err := client.TMDB.GetTVDetails(id, client.TMDBDefaultOptions)
	if err != nil {
		return nil, err
	}

	return ToMedia(&TVWrapper{tvShow}), nil
}

func (TVFetcher) Search(ctx context.Context, query string) ([]*Media, error) {
	tvShows, err := client.TMDB.GetSearchTVShow(query, client.TMDBDefaultOptions)
	if err != nil {
		return nil, err
	}

	var medias []*Media
	for _, tvShow := range tvShows.Results {
		var media TMDBMediaLight
		err := copier.Copy(&media, tvShow)
		if err != nil {
			return nil, err
		}

		media.Type = "tv"
		medias = append(medias, ToMedia(&media))
	}

	return medias, nil
}

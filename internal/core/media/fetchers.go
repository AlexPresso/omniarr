package media

import (
	"context"
	"github.com/Open-pi/gol"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	"omniarr/internal/client"
	"strconv"
)

type Fetcher interface {
	Fetch(ctx context.Context, id string) (*Media, error)
	Search(ctx context.Context, query string) ([]*Media, error)
}

type MovieFetcher struct{}
type TVFetcher struct{}
type MusicFetcher struct{}
type BookFetcher struct{}

func (MovieFetcher) Fetch(ctx context.Context, id string) (*Media, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	movie, err := client.TMDB.GetMovieDetails(intId, client.TMDBDefaultOptions)
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

func (TVFetcher) Fetch(ctx context.Context, id string) (*Media, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	tvShow, err := client.TMDB.GetTVDetails(intId, client.TMDBDefaultOptions)
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

func (BookFetcher) Fetch(ctx context.Context, id string) (*Media, error) {
	book, err := gol.GetWork(id)
	if err != nil {
		return nil, err
	}

	return ToMedia(&BookWrapper{&book}), nil
}

func (BookFetcher) Search(ctx context.Context, query string) ([]*Media, error) {
	books, err := gol.Search(gol.SearchUrl().All(query).Construct())
	if err != nil {
		return nil, err
	}

	var medias []*Media
	for _, book := range books.Path("docs").Data().([]interface{}) {
		var media OpenLibraryMediaLight
		if err := mapstructure.Decode(book, &media); err != nil {
			return nil, err
		}

		medias = append(medias, ToMedia(&media))
	}

	return medias, nil
}

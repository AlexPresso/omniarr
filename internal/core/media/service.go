package media

import (
	"context"
	tmdb "github.com/cyruzin/golang-tmdb"
	"omniarr/internal/client"
	"strconv"
)

func Search(ctx context.Context, query string) ([]Media, error) {
	videos, err := client.TMDBClient.GetSearchMovies(query, client.TMDBDefaultOptions)
	if err != nil {
		return nil, err
	}

	return mapToSearchResults(videos), nil
}

func GetDetails(ctx context.Context, id int) (*Details, error) {
	media, err := client.TMDBClient.GetMovieDetails(id, client.TMDBDefaultOptions)
	if err != nil {
		return nil, err
	}

	return mapToDetailsResult(media), nil
}

func mapToSearchResults(movies *tmdb.SearchMovies) []Media {
	res := make([]Media, 0, len(movies.Results))

	for _, movie := range movies.Results {
		res = append(res, Media{
			ID:            strconv.FormatInt(movie.ID, 10),
			Title:         movie.Title,
			OriginalTitle: movie.OriginalTitle,
			Description:   movie.Overview,
			Popularity:    movie.Popularity,
			ReleaseDate:   movie.ReleaseDate,
			Cover:         movie.PosterPath,
			Type:          "movie",
		})
	}

	return res
}

func mapToDetailsResult(details *tmdb.MovieDetails) *Details {
	return &Details{
		Media: &Media{
			ID:            strconv.FormatInt(details.ID, 10),
			Title:         details.Title,
			OriginalTitle: details.OriginalTitle,
			Description:   details.Overview,
			Popularity:    details.Popularity,
			ReleaseDate:   details.ReleaseDate,
			Cover:         details.PosterPath,
			Type:          "movie",
		},
	}
}

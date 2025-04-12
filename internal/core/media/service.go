package media

import (
	"context"
	tmdb "github.com/cyruzin/golang-tmdb"
	"omniarr/internal/config"
)

func Search(ctx context.Context, query string) ([]Result, error) {
	videos, err := config.TMDBClient.GetSearchMovies(query, nil)
	if err != nil {
		return nil, err
	}

	return mapToSearchResults(videos), nil
}

func mapToSearchResults(movies *tmdb.SearchMovies) []Result {
	res := make([]Result, 0, len(movies.Results))

	for _, movie := range movies.Results {
		res = append(res, Result{
			Title: movie.Title,
			Type:  "movie",
			Cover: "https://image.tmdb.org/t/p/original" + movie.PosterPath,
		})
	}

	return res
}

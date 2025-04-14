package client

import (
	"github.com/cyruzin/golang-tmdb"
	"log"
	"omniarr/internal/config"
)

var TMDB *tmdb.Client
var TMDBDefaultOptions map[string]string

func init() {
	client, err := tmdb.Init(config.AppConfig.TMDBAPIKey)
	if err != nil {
		log.Fatal(err)
	}

	client.SetClientAutoRetry()
	TMDBDefaultOptions = map[string]string{
		"language":           config.AppConfig.Language,
		"append_to_response": "alternative_titles",
	}
	TMDB = client

	log.Println("🔌 TMDB client initialized")
}

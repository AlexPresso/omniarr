package config

import (
	"github.com/cyruzin/golang-tmdb"
	"log"
)

var TMDBClient *tmdb.Client

func init() {
	client, err := tmdb.Init(AppConfig.TMDBAPIKey)
	if err != nil {
		log.Fatal(err)
	}

	client.SetClientAutoRetry()
	TMDBClient = client

	log.Println("🔌 TMDB client initialized")
}

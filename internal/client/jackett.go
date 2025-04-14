package client

import (
	"github.com/webtor-io/go-jackett"
	"log"
	"omniarr/internal/config"
)

var Jackett *jackett.Jackett

func init() {
	Jackett = jackett.NewJackett(&jackett.Settings{
		ApiURL: config.AppConfig.JackettURL,
		ApiKey: config.AppConfig.JackettApiKey,
	})

	log.Println("🔌 Jackett client initialized")
}

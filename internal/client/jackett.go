package client

import (
	"github.com/webtor-io/go-jackett"
	"log"
	"omniarr/internal/config"
)

var JackettClient *jackett.Jackett

func init() {
	JackettClient = jackett.NewJackett(&jackett.Settings{
		ApiURL: config.AppConfig.JackettURL,
		ApiKey: config.AppConfig.JackettApiKey,
	})

	log.Println("🔌 Jackett client initialized")
}

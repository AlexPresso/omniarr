package config

import (
	"github.com/webtor-io/go-jackett"
	"log"
)

var JackettClient *jackett.Jackett

func init() {
	JackettClient = jackett.NewJackett(&jackett.Settings{
		ApiURL: AppConfig.JackettURL,
		ApiKey: AppConfig.JackettApiKey,
	})

	log.Println("🔌 Jackett client initialized")
}

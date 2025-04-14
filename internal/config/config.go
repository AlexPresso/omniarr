package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port             string
	JackettURL       string
	JackettApiKey    string
	TMDBAPIKey       string
	Language         string
	TransmissionUser string
	TransmissionPass string
	TransmissionHost string
	TransmissionPort string
}

var AppConfig Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = Config{
		Port:             getEnv("PORT", "8080"),
		JackettURL:       getEnv("JACKETT_URL", "http://localhost:9117"),
		JackettApiKey:    getEnv("JACKETT_API_KEY", ""),
		TMDBAPIKey:       getEnv("TMDB_API_KEY", ""),
		Language:         getEnv("LANGUAGE", "en-US"),
		TransmissionUser: getEnv("TRANSMISSION_USER", ""),
		TransmissionPass: getEnv("TRANSMISSION_PASSWORD", ""),
		TransmissionHost: getEnv("TRANSMISSION_HOST", "localhost"),
		TransmissionPort: getEnv("TRANSMISSION_PORT", "9091"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	Port     string
}

var AppConfig Config

func loadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found (ok on production)")
	}

	AppConfig = Config{
		MongoURI: getEnv("MONGO_URI", ""),
		Port:     getEnv("PORT", "8080"),
	}

	if AppConfig.MongoURI == "" {
		log.Fatal("MONGO_URI not set")
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
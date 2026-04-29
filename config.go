package main

import ("log" "os")

type Config struct {
	MongoURI string
	Port	 string
}

var AppConfig Config

func loadConfig() {
	AppConfig = Config{
		MongoURI: getEnv("MONGO_URI", ""),
		Port: getEnv("PORT", "8080"),
	}

	if AppConfig.MongoURI == "" {
		log.Fatal("Mongo URI failed")
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}

	return val
}
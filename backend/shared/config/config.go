package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI          string
	JWTSecret         string
	FirebaseConfigPath string
	Port              string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		MongoURI:          getEnvWithDefault("MONGO_URI", "mongodb://localhost:27017"),
		JWTSecret:         getEnvWithDefault("JWT_SECRET", "your-super-secret-jwt-key"),
		FirebaseConfigPath: getEnvWithDefault("FIREBASE_CONFIG_PATH", "firebase-config.json"),
		Port:              getEnvWithDefault("PORT", "8080"),
	}
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
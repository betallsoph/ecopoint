package config

import (
	"os"
	"strings"
)

type Config struct {
	Environment string
	Port        string
	UserServiceURL string
	BookingServiceURL string
	NotificationServiceURL string
	CORSOrigins []string
	FirebaseProjectID string
}

func Load() *Config {
	return &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		Port:        getEnv("PORT", "8080"),
		UserServiceURL: getEnv("USER_SERVICE_URL", "http://localhost:4000"),
		BookingServiceURL: getEnv("BOOKING_SERVICE_URL", "http://localhost:50051"),
		NotificationServiceURL: getEnv("NOTIFICATION_SERVICE_URL", "http://localhost:50052"),
		CORSOrigins: strings.Split(getEnv("CORS_ORIGINS", "http://localhost:3000,http://localhost:3001"), ","),
		FirebaseProjectID: getEnv("FIREBASE_PROJECT_ID", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
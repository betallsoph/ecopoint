package config

import (
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    MongoURI        string
    MongoDBName     string
    OrdersTTLMinutes int
}

func Load() *Config {
    _ = godotenv.Load()
    ttl := 60
    if v := os.Getenv("ORDERS_TTL_MINUTES"); v != "" {
        if n, err := strconv.Atoi(v); err == nil && n > 0 {
            ttl = n
        }
    }
    dbName := os.Getenv("MONGO_DB_NAME")
    if dbName == "" {
        dbName = "ecopoint"
    }
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        uri = "mongodb://localhost:27017/" + dbName
    }
    return &Config{
        MongoURI: uri,
        MongoDBName: dbName,
        OrdersTTLMinutes: ttl,
    }
}



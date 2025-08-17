package main

import (
    "context"
    "fmt"
    "log"

    "ecopoint/collecting_service/internal/config"
    "ecopoint/collecting_service/internal/repository"
)

func main() {
    cfg := config.Load()
    ctx := context.Background()

    repo, err := repository.NewMongoRepo(ctx, cfg.MongoURI, cfg.MongoDBName)
    if err != nil { log.Fatalf("mongo connect error: %v", err) }
    defer repo.Close(ctx)

    if err := repo.InitIndexes(ctx); err != nil {
        log.Fatalf("init indexes error: %v", err)
    }

    fmt.Println("Mongo connected and indexes ensured.")
}



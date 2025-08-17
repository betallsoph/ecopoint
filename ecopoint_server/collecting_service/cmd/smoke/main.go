package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "time"

    "ecopoint/collecting_service/internal/config"
    "ecopoint/collecting_service/internal/models"
    "ecopoint/collecting_service/internal/repository"
    "ecopoint/collecting_service/internal/service"
)

func pp(label string, v any) {
    b, _ := json.MarshalIndent(v, "", "  ")
    fmt.Printf("\n=== %s ===\n%s\n", label, string(b))
}

func main() {
    ctx := context.Background()
    cfg := config.Load()

    repo, err := repository.NewMongoRepo(ctx, cfg.MongoURI, cfg.MongoDBName)
    if err != nil {
        log.Fatalf("mongo connect error: %v", err)
    }
    defer repo.Close(ctx)
    if err := repo.InitIndexes(ctx); err != nil {
        log.Fatalf("init indexes error: %v", err)
    }

    svc := service.NewService(repo)

    // 1) Create order
    id := fmt.Sprintf("smoke_%d", time.Now().UnixNano())
    order, err := svc.CreateOrder(service.CreateOrderInput{
        ID:         id,
        CustomerID: "user_demo_1",
        Address:    models.Address{FullText: "123 Demo St, HCMC", Lat: 10.775, Lng: 106.700},
        CustomerSnapshot: models.CustomerSnapshot{DisplayName: "Demo User", Phone: "0909000000"},
        Items: []models.WasteItem{{Type: "plastic", Weight: 1.2}, {Type: "paper", Weight: 0.8}},
        TotalWeight:    2.0,
        EstimatedPrice: 50000,
        Note:           "smoke test",
    })
    if err != nil { log.Fatalf("CreateOrder error: %v", err) }
    pp("CreateOrder", order)

    // 2) List available
    list, err := svc.ListAvailable(10)
    if err != nil { log.Fatalf("ListAvailable error: %v", err) }
    fmt.Printf("Available count: %d\n", len(list))

    // 3) Accept order by collector
    accepted, err := svc.AcceptOrder(id, "collector_demo")
    if err != nil { log.Fatalf("AcceptOrder error: %v", err) }
    pp("AcceptOrder", accepted)

    // 4) Double accept should fail
    if _, err := svc.AcceptOrder(id, "collector_other"); err != nil {
        fmt.Println("Double accept blocked as expected:", err)
    } else {
        log.Fatal("Double accept unexpectedly succeeded")
    }

    // 5) List my active orders
    actives, err := svc.ListMyActiveOrders("collector_demo")
    if err != nil { log.Fatalf("ListMyActiveOrders error: %v", err) }
    fmt.Printf("Active orders for collector_demo: %d\n", len(actives))

    // 6) Update to on_way then complete
    onway, err := svc.UpdateStatus(id, models.StatusOnWay, "collector_demo")
    if err != nil { log.Fatalf("Update to on_way error: %v", err) }
    pp("Update to on_way", onway)

    done, err := svc.UpdateStatus(id, models.StatusComplete, "collector_demo")
    if err != nil { log.Fatalf("Update to complete error: %v", err) }
    pp("Update to complete", done)

    // 7) Get order final
    got, err := svc.GetOrder(id)
    if err != nil { log.Fatalf("GetOrder error: %v", err) }
    pp("GetOrder final", got)

    fmt.Println("Smoke test finished OK.")
}



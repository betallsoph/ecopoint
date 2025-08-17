package service

import (
    "testing"
    "time"

    "ecopoint/collecting_service/internal/models"
)

func TestCreateAndList(t *testing.T) {
    repo := NewInMemoryRepo()
    svc := NewService(repo)

    // Initially empty
    list, err := svc.ListAvailable(10)
    if err != nil || len(list) != 0 {
        t.Fatalf("expected empty list, got %v err %v", len(list), err)
    }

    // Create
    order, err := svc.CreateOrder(CreateOrderInput{
        ID:         "o1",
        CustomerID: "u1",
        Address: models.Address{FullText: "A", Lat: 1, Lng: 2},
        CustomerSnapshot: models.CustomerSnapshot{DisplayName: "Name", Phone: "0909"},
        Items: []models.WasteItem{{Type: "plastic", Weight: 1.2}},
        TotalWeight: 1.2,
        EstimatedPrice: 10000,
        Note: "n",
    })
    if err != nil {
        t.Fatalf("create error: %v", err)
    }
    if order.Status != models.StatusCreated {
        t.Fatalf("status expected created, got %s", order.Status)
    }

    list, err = svc.ListAvailable(10)
    if err != nil || len(list) != 1 {
        t.Fatalf("expected 1 available order, got %v err %v", len(list), err)
    }
}

func TestAcceptAndUpdateStatus(t *testing.T) {
    repo := NewInMemoryRepo()
    svc := NewService(repo)

    // Create
    _, _ = svc.CreateOrder(CreateOrderInput{
        ID:         "o2",
        CustomerID: "u1",
        Address: models.Address{FullText: "A", Lat: 1, Lng: 2},
        CustomerSnapshot: models.CustomerSnapshot{DisplayName: "Name", Phone: "0909"},
        Items: []models.WasteItem{{Type: "paper", Weight: 2}},
        TotalWeight: 2,
        EstimatedPrice: 20000,
        Note: "n2",
    })

    // Accept
    accepted, err := svc.AcceptOrder("o2", "collector-1")
    if err != nil {
        t.Fatalf("accept error: %v", err)
    }
    if accepted.Status != models.StatusAccepted {
        t.Fatalf("expected accepted, got %s", accepted.Status)
    }

    // Double accept should fail
    if _, err := svc.AcceptOrder("o2", "collector-2"); err == nil {
        t.Fatalf("expected error on double accept")
    }

    // Move to on_way, then complete
    o, err := svc.UpdateStatus("o2", models.StatusOnWay, "collector-1")
    if err != nil || o.Status != models.StatusOnWay {
        t.Fatalf("to on_way failed: %v status %s", err, o.Status)
    }
    o, err = svc.UpdateStatus("o2", models.StatusComplete, "collector-1")
    if err != nil || o.Status != models.StatusComplete {
        t.Fatalf("to complete failed: %v status %s", err, o.Status)
    }

    // Invalid transition after complete
    if _, err := svc.UpdateStatus("o2", models.StatusOnWay, "collector-1"); err == nil {
        t.Fatalf("expected invalid transition error")
    }
}

func TestCancelRules(t *testing.T) {
    repo := NewInMemoryRepo()
    svc := NewService(repo)

    // Customer can cancel when created
    _, _ = svc.CreateOrder(CreateOrderInput{ID: "oc1", CustomerID: "u1"})
    if _, err := svc.CancelOrderByCustomer("oc1", "change of mind"); err != nil {
        t.Fatalf("customer cancel failed: %v", err)
    }

    // After accepted, customer cannot cancel
    _, _ = svc.CreateOrder(CreateOrderInput{ID: "oc2", CustomerID: "u1"})
    _, _ = svc.AcceptOrder("oc2", "c1")
    if _, err := svc.CancelOrderByCustomer("oc2", "late"); err == nil {
        t.Fatalf("expected error: customer cancel after accepted")
    }

    // Collector can cancel when accepted
    if _, err := svc.CancelOrderByCollector("oc2", "c1", "busy"); err != nil {
        t.Fatalf("collector cancel failed: %v", err)
    }
}

func TestOneActiveOrderPerCollector(t *testing.T) {
    repo := NewInMemoryRepo()
    svc := NewService(repo)

    _, _ = svc.CreateOrder(CreateOrderInput{ID: "oa1", CustomerID: "u1"})
    _, _ = svc.CreateOrder(CreateOrderInput{ID: "oa2", CustomerID: "u2"})

    if _, err := svc.AcceptOrder("oa1", "collector-1"); err != nil {
        t.Fatalf("accept oa1 failed: %v", err)
    }
    // Should not accept another while active
    if _, err := svc.AcceptOrder("oa2", "collector-1"); err == nil {
        t.Fatalf("expected error: collector already has active order")
    }
}

func TestPricingAndETA(t *testing.T) {
    svc := NewService(NewInMemoryRepo())
    price, eta := svc.ComputePriceAndETA(10000, 2000, 3000, 30, 2.5, 5)
    if price <= 0 || eta <= 0 {
        t.Fatalf("unexpected price/eta %v %v", price, eta)
    }
}

func TestListAvailableOrdersNearAndExpire(t *testing.T) {
    repo := NewInMemoryRepo()
    svc := NewService(repo)

    // Create two orders with different locations
    _, _ = svc.CreateOrder(CreateOrderInput{
        ID: "n1", CustomerID: "u1",
        Address: models.Address{FullText: "A", Lat: 10.76, Lng: 106.66},
    })
    _, _ = svc.CreateOrder(CreateOrderInput{
        ID: "n2", CustomerID: "u2",
        Address: models.Address{FullText: "B", Lat: 10.80, Lng: 106.70},
    })

    // Near search around (10.77, 106.67) within 5km
    res, err := svc.ListAvailableOrdersNear(10.77, 106.67, 5, 10)
    if err != nil || len(res) == 0 {
        t.Fatalf("near search failed: %v len=%d", err, len(res))
    }

    // Expire: set created_at back 2h for n1 then run TTL=60
    o1, _ := repo.Get("n1")
    o1.CreatedAt = o1.CreatedAt.Add(-2 * time.Hour) // -2h
    _ = repo.Update(o1)
    expired, err := svc.AutoExpireCreatedOrders(60)
    if err != nil || expired < 1 {
        t.Fatalf("expire failed: %v expired=%d", err, expired)
    }
}

func TestNewAPIs(t *testing.T) {
    repo := NewInMemoryRepo()
    svc := NewService(repo)

    // Create 3 orders for customer u9
    _, _ = svc.CreateOrder(CreateOrderInput{ID: "m1", CustomerID: "u9"})
    _, _ = svc.CreateOrder(CreateOrderInput{ID: "m2", CustomerID: "u9"})
    _, _ = svc.CreateOrder(CreateOrderInput{ID: "m3", CustomerID: "u9"})

    // ListMyOrders page 1 size 2
    list, err := svc.ListMyOrders("u9", 1, 2)
    if err != nil || len(list) != 2 {
        t.Fatalf("ListMyOrders failed: %v len=%d", err, len(list))
    }

    if _, err := svc.AcceptOrder("m1", "c7"); err != nil {
        t.Fatalf("accept m1 failed: %v", err)
    }
    actives, err := svc.ListMyActiveOrders("c7")
    if err != nil || len(actives) != 1 {
        t.Fatalf("ListMyActiveOrders failed: %v len=%d", err, len(actives))
    }

    o, err := svc.GetOrder("m2")
    if err != nil || o.ID != "m2" {
        t.Fatalf("GetOrder failed: %v id=%s", err, o.ID)
    }
}



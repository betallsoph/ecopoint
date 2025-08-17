package service

import (
    "errors"
    "fmt"
    "math"
    "sort"
    "time"

    "ecopoint/collecting_service/internal/models"
)

// Repository abstracts order storage (in-memory for now)
type Repository interface {
    Create(order *models.Order) error
    Get(id string) (*models.Order, error)
    ListAvailable(limit int) ([]*models.Order, error)
    AtomicAccept(id string, collectorID string) (*models.Order, error)
    FindActiveOrderByCollector(collectorID string) (*models.Order, error)
    Update(order *models.Order) error
    ListAll() ([]*models.Order, error)
    // Optional optimized queries for convenience
    ListByCustomer(customerID string, page, size int) ([]*models.Order, error)
    ListActiveByCollector(collectorID string) ([]*models.Order, error)
}

// InMemoryRepo is a simple in-memory implementation for tests
type InMemoryRepo struct {
    store map[string]*models.Order
}

func NewInMemoryRepo() *InMemoryRepo {
    return &InMemoryRepo{store: map[string]*models.Order{}}
}

func (r *InMemoryRepo) Create(order *models.Order) error {
    if _, ok := r.store[order.ID]; ok {
        return fmt.Errorf("order exists")
    }
    r.store[order.ID] = order
    return nil
}

func (r *InMemoryRepo) Get(id string) (*models.Order, error) {
    o, ok := r.store[id]
    if !ok {
        return nil, fmt.Errorf("not found")
    }
    return o, nil
}

func (r *InMemoryRepo) ListAvailable(limit int) ([]*models.Order, error) {
    res := make([]*models.Order, 0, limit)
    for _, o := range r.store {
        if o.Status == models.StatusCreated {
            res = append(res, o)
            if len(res) >= limit {
                break
            }
        }
    }
    return res, nil
}

func (r *InMemoryRepo) AtomicAccept(id string, collectorID string) (*models.Order, error) {
    o, ok := r.store[id]
    if !ok {
        return nil, fmt.Errorf("not found")
    }
    if o.Status != models.StatusCreated {
        return nil, fmt.Errorf("already taken")
    }
    now := time.Now()
    o.Status = models.StatusAccepted
    o.AcceptedBy = &collectorID
    o.AcceptedAt = &now
    o.UpdatedAt = now
    o.Version++
    return o, nil
}

func (r *InMemoryRepo) Update(order *models.Order) error {
    if _, ok := r.store[order.ID]; !ok {
        return fmt.Errorf("not found")
    }
    r.store[order.ID] = order
    return nil
}

func (r *InMemoryRepo) FindActiveOrderByCollector(collectorID string) (*models.Order, error) {
    for _, o := range r.store {
        if o.AcceptedBy != nil && *o.AcceptedBy == collectorID && (o.Status == models.StatusAccepted || o.Status == models.StatusOnWay) {
            return o, nil
        }
    }
    return nil, fmt.Errorf("not found")
}

func (r *InMemoryRepo) ListAll() ([]*models.Order, error) {
    res := make([]*models.Order, 0, len(r.store))
    for _, o := range r.store {
        res = append(res, o)
    }
    return res, nil
}

func (r *InMemoryRepo) ListByCustomer(customerID string, page, size int) ([]*models.Order, error) {
    if page < 1 { page = 1 }
    if size <= 0 { size = 20 }
    start := (page-1) * size
    // naive: collect all then paginate
    all := make([]*models.Order, 0)
    for _, o := range r.store {
        if o.CustomerID == customerID {
            all = append(all, o)
        }
    }
    // sort by CreatedAt desc
    sort.Slice(all, func(i, j int) bool { return all[i].CreatedAt.After(all[j].CreatedAt) })
    end := start + size
    if start >= len(all) { return []*models.Order{}, nil }
    if end > len(all) { end = len(all) }
    return all[start:end], nil
}

func (r *InMemoryRepo) ListActiveByCollector(collectorID string) ([]*models.Order, error) {
    res := make([]*models.Order, 0)
    for _, o := range r.store {
        if o.AcceptedBy != nil && *o.AcceptedBy == collectorID && (o.Status == models.StatusAccepted || o.Status == models.StatusOnWay) {
            res = append(res, o)
        }
    }
    return res, nil
}

// Service contains business logic
type Service struct {
    repo Repository
}

func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}

type CreateOrderInput struct {
    ID               string
    CustomerID       string
    Address          models.Address
    CustomerSnapshot models.CustomerSnapshot
    Items            []models.WasteItem
    TotalWeight      float64
    EstimatedPrice   float64
    Note             string
}

func (s *Service) CreateOrder(in CreateOrderInput) (*models.Order, error) {
    now := time.Now()
    order := &models.Order{
        ID:                  in.ID,
        CustomerID:          in.CustomerID,
        Status:              models.StatusCreated,
        PickAddressSnapshot: in.Address,
        CustomerSnapshot:    in.CustomerSnapshot,
        Items:               in.Items,
        TotalWeight:         in.TotalWeight,
        EstimatedPrice:      in.EstimatedPrice,
        Note:                in.Note,
        CreatedAt:           now,
        UpdatedAt:           now,
        Version:             1,
    }
    if err := s.repo.Create(order); err != nil {
        return nil, err
    }
    return order, nil
}

func (s *Service) ListAvailable(limit int) ([]*models.Order, error) {
    return s.repo.ListAvailable(limit)
}

func (s *Service) AcceptOrder(orderID string, collectorID string) (*models.Order, error) {
    // Rule: one active order per collector
    if _, err := s.repo.FindActiveOrderByCollector(collectorID); err == nil {
        return nil, errors.New("collector already has an active order")
    }
    return s.repo.AtomicAccept(orderID, collectorID)
}

func (s *Service) UpdateStatus(orderID string, next models.OrderStatus, collectorID string) (*models.Order, error) {
    o, err := s.repo.Get(orderID)
    if err != nil {
        return nil, err
    }
    if o.AcceptedBy == nil || *o.AcceptedBy != collectorID {
        return nil, errors.New("not owner")
    }
    if !o.CanTransition(next) {
        return nil, models.ErrInvalidStatusTransition
    }
    now := time.Now()
    o.Status = next
    if next == models.StatusComplete {
        o.CompletedAt = &now
    }
    o.UpdatedAt = now
    o.Version++
    if err := s.repo.Update(o); err != nil {
        return nil, err
    }
    return o, nil
}

// New APIs
func (s *Service) GetOrder(orderID string) (*models.Order, error) {
    return s.repo.Get(orderID)
}

func (s *Service) ListMyActiveOrders(collectorID string) ([]*models.Order, error) {
    return s.repo.ListActiveByCollector(collectorID)
}

func (s *Service) ListMyOrders(customerID string, page, size int) ([]*models.Order, error) {
    return s.repo.ListByCustomer(customerID, page, size)
}

// 1) Cancel by customer: only when not yet accepted
func (s *Service) CancelOrderByCustomer(orderID string, reason string) (*models.Order, error) {
    o, err := s.repo.Get(orderID)
    if err != nil {
        return nil, err
    }
    if o.Status != models.StatusCreated {
        return nil, errors.New("cannot cancel after accepted")
    }
    now := time.Now()
    o.Status = models.StatusCancelled
    o.CancelSide = models.CancelByCustomer
    o.CancelReason = reason
    o.UpdatedAt = now
    o.Version++
    if err := s.repo.Update(o); err != nil {
        return nil, err
    }
    return o, nil
}

// 2) Cancel by collector: allowed when accepted/on_way
func (s *Service) CancelOrderByCollector(orderID string, collectorID string, reason string) (*models.Order, error) {
    o, err := s.repo.Get(orderID)
    if err != nil {
        return nil, err
    }
    if o.AcceptedBy == nil || *o.AcceptedBy != collectorID {
        return nil, errors.New("not owner")
    }
    if !(o.Status == models.StatusAccepted || o.Status == models.StatusOnWay) {
        return nil, errors.New("cannot cancel at this status")
    }
    now := time.Now()
    o.Status = models.StatusCancelled
    o.CancelSide = models.CancelByCollector
    o.CancelReason = reason
    o.UpdatedAt = now
    o.Version++
    if err := s.repo.Update(o); err != nil {
        return nil, err
    }
    return o, nil
}

// 3) Pricing + ETA (simple): base + weight_factor*kg + distance_factor*km; ETA = distance/avg_speed
func (s *Service) ComputePriceAndETA(base, weightFactor, distanceFactor, avgSpeedKmH float64, weightKg, distanceKm float64) (price float64, etaMinutes int) {
    price = base + weightFactor*weightKg + distanceFactor*distanceKm
    if avgSpeedKmH <= 0 {
        etaMinutes = 0
    } else {
        etaMinutes = int((distanceKm/avgSpeedKmH)*60 + 0.5)
    }
    if price < 0 {
        price = 0
    }
    return
}

// 4) ListAvailableOrdersNear with Haversine filter (in-memory)
func (s *Service) ListAvailableOrdersNear(lat, lng, radiusKm float64, limit int) ([]*models.Order, error) {
    all, err := s.repo.ListAll()
    if err != nil {
        return nil, err
    }
    res := make([]*models.Order, 0, limit)
    type scored struct {
        o   *models.Order
        km  float64
    }
    scores := make([]scored, 0)
    for _, o := range all {
        if o.Status != models.StatusCreated {
            continue
        }
        d := haversineKm(lat, lng, o.PickAddressSnapshot.Lat, o.PickAddressSnapshot.Lng)
        if d <= radiusKm {
            scores = append(scores, scored{o: o, km: d})
        }
    }
    // sort by distance then created_at desc
    sort.Slice(scores, func(i, j int) bool {
        if scores[i].km == scores[j].km {
            return scores[i].o.CreatedAt.After(scores[j].o.CreatedAt)
        }
        return scores[i].km < scores[j].km
    })
    for _, sc := range scores {
        res = append(res, sc.o)
        if len(res) >= limit {
            break
        }
    }
    return res, nil
}

// 5) Auto-expire created orders older than ttl minutes
func (s *Service) AutoExpireCreatedOrders(ttlMinutes int) (expired int, err error) {
    all, err := s.repo.ListAll()
    if err != nil {
        return 0, err
    }
    cutoff := time.Now().Add(-time.Duration(ttlMinutes) * time.Minute)
    for _, o := range all {
        if o.Status == models.StatusCreated && o.CreatedAt.Before(cutoff) {
            o.Status = models.StatusCancelled
            o.CancelSide = models.CancelBySystem
            o.CancelReason = "expired"
            o.UpdatedAt = time.Now()
            o.Version++
            if err := s.repo.Update(o); err != nil {
                return expired, err
            }
            expired++
        }
    }
    return expired, nil
}

// Helpers
func haversineKm(lat1, lon1, lat2, lon2 float64) float64 {
    const R = 6371.0
    toRad := func(d float64) float64 { return d * math.Pi / 180 }
    dLat := toRad(lat2 - lat1)
    dLon := toRad(lon2 - lon1)
    a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(toRad(lat1))*math.Cos(toRad(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)
    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
    return R * c
}



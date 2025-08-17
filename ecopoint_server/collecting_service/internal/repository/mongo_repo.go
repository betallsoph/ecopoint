package repository

import (
    "context"
    "errors"
    "time"

    "ecopoint/collecting_service/internal/models"
    svc "ecopoint/collecting_service/internal/service"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
    client    *mongo.Client
    db        *mongo.Database
    ordersCol *mongo.Collection
}

func NewMongoRepo(ctx context.Context, uri string, dbName string) (*MongoRepo, error) {
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        return nil, err
    }
    if err := client.Ping(ctx, nil); err != nil {
        return nil, err
    }
    db := client.Database(dbName)
    repo := &MongoRepo{
        client:    client,
        db:        db,
        ordersCol: db.Collection("orders"),
    }
    return repo, nil
}

func (r *MongoRepo) Close(ctx context.Context) error {
    return r.client.Disconnect(ctx)
}

// InitIndexes creates recommended indexes for performance and TTL
func (r *MongoRepo) InitIndexes(ctx context.Context) error {
    // unique id
    _, err := r.ordersCol.Indexes().CreateOne(ctx, mongo.IndexModel{
        Keys:    bson.D{{Key: "id", Value: 1}},
        Options: options.Index().SetUnique(true),
    })
    if err != nil {
        return err
    }
    // status + created_at
    _, err = r.ordersCol.Indexes().CreateOne(ctx, mongo.IndexModel{
        Keys: bson.D{{Key: "status", Value: 1}, {Key: "created_at", Value: -1}},
    })
    if err != nil {
        return err
    }
    // accepted_by + status
    _, err = r.ordersCol.Indexes().CreateOne(ctx, mongo.IndexModel{
        Keys: bson.D{{Key: "accepted_by", Value: 1}, {Key: "status", Value: 1}},
    })
    if err != nil {
        return err
    }
    // TTL on expire_at
    _, err = r.ordersCol.Indexes().CreateOne(ctx, mongo.IndexModel{
        Keys:    bson.D{{Key: "expire_at", Value: 1}},
        Options: options.Index().SetExpireAfterSeconds(0),
    })
    if err != nil {
        return err
    }
    // Geo index (optional) if using loc
    _, _ = r.ordersCol.Indexes().CreateOne(ctx, mongo.IndexModel{
        Keys: bson.D{{Key: "loc", Value: "2dsphere"}},
    })
    return nil
}

// Helpers to build the Mongo document from domain model
func orderToDoc(o *models.Order) bson.M {
    doc := bson.M{
        "id":                   o.ID,
        "customer_id":          o.CustomerID,
        "status":               o.Status,
        "pick_address_snapshot": o.PickAddressSnapshot,
        "customer_snapshot":     o.CustomerSnapshot,
        "items":                o.Items,
        "total_weight":         o.TotalWeight,
        "estimated_price":      o.EstimatedPrice,
        "distance_km":          o.DistanceKm,
        "eta_minutes":          o.EtaMinutes,
        "note":                 o.Note,
        "created_at":           o.CreatedAt,
        "updated_at":           o.UpdatedAt,
        "version":              o.Version,
    }
    if o.AcceptedBy != nil {
        doc["accepted_by"] = *o.AcceptedBy
    }
    if o.AcceptedAt != nil {
        doc["accepted_at"] = *o.AcceptedAt
    }
    if o.CompletedAt != nil {
        doc["completed_at"] = *o.CompletedAt
    }
    if o.CancelReason != "" {
        doc["cancel_reason"] = o.CancelReason
        doc["cancel_side"] = o.CancelSide
    }
    // GeoJSON point for near queries
    doc["loc"] = bson.M{"type": "Point", "coordinates": []float64{o.PickAddressSnapshot.Lng, o.PickAddressSnapshot.Lat}}
    return doc
}

func docToOrder(doc *bson.M) *models.Order {
    if doc == nil {
        return nil
    }
    var o models.Order
    b, _ := bson.Marshal(doc)
    _ = bson.Unmarshal(b, &o)
    return &o
}

// Implement service.Repository
func (r *MongoRepo) Create(order *models.Order) error {
    // set expire_at for created orders (TTL), e.g., 60 minutes default here; can be moved to config
    expireAt := order.CreatedAt.Add(60 * time.Minute)
    doc := orderToDoc(order)
    if order.Status == models.StatusCreated {
        doc["expire_at"] = expireAt
    }
    _, err := r.ordersCol.InsertOne(context.Background(), doc)
    return err
}

func (r *MongoRepo) Get(id string) (*models.Order, error) {
    var m bson.M
    err := r.ordersCol.FindOne(context.Background(), bson.M{"id": id}).Decode(&m)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) { return nil, errors.New("not found") }
        return nil, err
    }
    return docToOrder(&m), nil
}

func (r *MongoRepo) ListAvailable(limit int) ([]*models.Order, error) {
    opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(int64(limit))
    cursor, err := r.ordersCol.Find(context.Background(), bson.M{"status": models.StatusCreated}, opts)
    if err != nil { return nil, err }
    defer cursor.Close(context.Background())
    var res []*models.Order
    for cursor.Next(context.Background()) {
        var m bson.M
        if err := cursor.Decode(&m); err != nil { return nil, err }
        res = append(res, docToOrder(&m))
    }
    return res, cursor.Err()
}

func (r *MongoRepo) AtomicAccept(id string, collectorID string) (*models.Order, error) {
    now := time.Now()
    filter := bson.M{"id": id, "status": models.StatusCreated}
    update := bson.M{
        "$set": bson.M{
            "status":      models.StatusAccepted,
            "accepted_by": collectorID,
            "accepted_at": now,
            "updated_at":  now,
        },
        "$inc": bson.M{"version": 1},
        "$unset": bson.M{"expire_at": ""},
    }
    opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
    var m bson.M
    err := r.ordersCol.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&m)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) { return nil, errors.New("already taken or not found") }
        return nil, err
    }
    return docToOrder(&m), nil
}

func (r *MongoRepo) Update(order *models.Order) error {
    doc := orderToDoc(order)
    _, err := r.ordersCol.UpdateOne(context.Background(), bson.M{"id": order.ID}, bson.M{"$set": doc})
    return err
}

func (r *MongoRepo) FindActiveOrderByCollector(collectorID string) (*models.Order, error) {
    filter := bson.M{"accepted_by": collectorID, "status": bson.M{"$in": []models.OrderStatus{models.StatusAccepted, models.StatusOnWay}}}
    var m bson.M
    err := r.ordersCol.FindOne(context.Background(), filter).Decode(&m)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) { return nil, errors.New("not found") }
        return nil, err
    }
    return docToOrder(&m), nil
}

func (r *MongoRepo) ListAll() ([]*models.Order, error) {
    cursor, err := r.ordersCol.Find(context.Background(), bson.M{})
    if err != nil { return nil, err }
    defer cursor.Close(context.Background())
    var res []*models.Order
    for cursor.Next(context.Background()) {
        var m bson.M
        if err := cursor.Decode(&m); err != nil { return nil, err }
        res = append(res, docToOrder(&m))
    }
    return res, cursor.Err()
}

func (r *MongoRepo) ListByCustomer(customerID string, page, size int) ([]*models.Order, error) {
    if page < 1 { page = 1 }
    if size <= 0 { size = 20 }
    skip := int64((page-1) * size)
    limit := int64(size)
    opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetSkip(skip).SetLimit(limit)
    cursor, err := r.ordersCol.Find(context.Background(), bson.M{"customer_id": customerID}, opts)
    if err != nil { return nil, err }
    defer cursor.Close(context.Background())
    var res []*models.Order
    for cursor.Next(context.Background()) {
        var m bson.M
        if err := cursor.Decode(&m); err != nil { return nil, err }
        res = append(res, docToOrder(&m))
    }
    return res, cursor.Err()
}

func (r *MongoRepo) ListActiveByCollector(collectorID string) ([]*models.Order, error) {
    filter := bson.M{"accepted_by": collectorID, "status": bson.M{"$in": []models.OrderStatus{models.StatusAccepted, models.StatusOnWay}}}
    cursor, err := r.ordersCol.Find(context.Background(), filter)
    if err != nil { return nil, err }
    defer cursor.Close(context.Background())
    var res []*models.Order
    for cursor.Next(context.Background()) {
        var m bson.M
        if err := cursor.Decode(&m); err != nil { return nil, err }
        res = append(res, docToOrder(&m))
    }
    return res, cursor.Err()
}

// Ensure MongoRepo implements Repository
var _ svc.Repository = (*MongoRepo)(nil)



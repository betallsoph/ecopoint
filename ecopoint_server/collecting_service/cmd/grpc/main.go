package main

import (
	"context"
	"log"
	"net"
	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "ecopoint/collecting_service/pb"
	"ecopoint/collecting_service/internal/config"
	"ecopoint/collecting_service/internal/models"
	"ecopoint/collecting_service/internal/repository"
	"ecopoint/collecting_service/internal/service"
)

type server struct {
	pb.UnimplementedCollectingServiceServer
	svc *service.Service
}


func (s *server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {

	o, err := s.svc.CreateOrder(service.CreateOrderInput{
		ID:               uuid.NewString(),
		CustomerID:       req.CustomerId,
		Address:          addressPbToModel(req.PickAddress),
		CustomerSnapshot: customerPbToModel(req.CustomerSnapshot),
		Items:            itemsPbToModel(req.Items),
		TotalWeight:      req.TotalWeight,
		EstimatedPrice:   req.EstimatedPrice,
		Note:             req.Note,
	})
	if err != nil {
		return nil, err
	}
	return orderModelToPb(o), nil
}

func (s *server) ListAvailableOrders(ctx context.Context, req *pb.ListAvailableOrdersRequest) (*pb.ListAvailableOrdersResponse, error) {
    limit := int(req.Limit)
    if limit <= 0 { limit = 20 }
    list, err := s.svc.ListAvailable(limit)
    if err != nil { return nil, err }
    res := &pb.ListAvailableOrdersResponse{}
    for _, o := range list { res.Orders = append(res.Orders, orderModelToPb(o)) }
    return res, nil
}

func (s *server) AcceptOrder(ctx context.Context, req *pb.AcceptOrderRequest) (*pb.Order, error) {
    o, err := s.svc.AcceptOrder(req.OrderId, req.CollectorId)
    if err != nil { return nil, err }
    return orderModelToPb(o), nil
}

func (s *server) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.Order, error) {
    o, err := s.svc.UpdateStatus(req.OrderId, statusFromString(req.Status), req.CollectorId)
    if err != nil { return nil, err }
    return orderModelToPb(o), nil
}

func (s *server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
    o, err := s.svc.GetOrder(req.OrderId)
    if err != nil { return nil, err }
    return orderModelToPb(o), nil
}

func (s *server) ListMyActiveOrders(ctx context.Context, req *pb.ListMyActiveOrdersRequest) (*pb.ListOrdersResponse, error) {
    list, err := s.svc.ListMyActiveOrders(req.CollectorId)
    if err != nil { return nil, err }
    res := &pb.ListOrdersResponse{}
    for _, o := range list { res.Orders = append(res.Orders, orderModelToPb(o)) }
    return res, nil
}

func (s *server) ListMyOrders(ctx context.Context, req *pb.ListMyOrdersRequest) (*pb.ListOrdersResponse, error) {
    list, err := s.svc.ListMyOrders(req.CustomerId, int(req.Page), int(req.Size))
    if err != nil { return nil, err }
    res := &pb.ListOrdersResponse{}
    for _, o := range list { res.Orders = append(res.Orders, orderModelToPb(o)) }
    return res, nil
}

func main(){
    cfg := config.Load()
    ctx := context.Background()
    repo, err := repository.NewMongoRepo(ctx, cfg.MongoURI, cfg.MongoDBName)
    if err != nil { log.Fatalf("mongo: %v", err) }
    defer repo.Close(ctx)
    _ = repo.InitIndexes(ctx)

    s := &server{ svc: service.NewService(repo) }

    grpcServer := grpc.NewServer()
    pb.RegisterCollectingServiceServer(grpcServer, s)
    reflection.Register(grpcServer)

    lis, err := net.Listen("tcp", ":50052")
    if err != nil { log.Fatalf("listen: %v", err) }
    log.Println("Collecting gRPC listening on :50052")
    if err := grpcServer.Serve(lis); err != nil { log.Fatal(err) }
}

// mapping helpers
func orderModelToPb(o *models.Order) *pb.Order {
	return &pb.Order{
		Id:        o.ID,
		CustomerId: o.CustomerID,
		Status:    string(o.Status),
		AcceptedBy: valueOrEmpty(o.AcceptedBy),
		PickAddressSnapshot: &pb.Address{
			FullText: o.PickAddressSnapshot.FullText,
			Lat:      o.PickAddressSnapshot.Lat,
			Lng:      o.PickAddressSnapshot.Lng,
		},
		CustomerSnapshot: &pb.CustomerSnapshot{
			DisplayName: o.CustomerSnapshot.DisplayName,
			Phone:       o.CustomerSnapshot.Phone,
		},
		TotalWeight:    o.TotalWeight,
		EstimatedPrice: o.EstimatedPrice,
		DistanceKm:     o.DistanceKm,
		EtaMinutes:     int32(o.EtaMinutes),
		Note:           o.Note,
		Version:        o.Version,
	}
}

// lightweight helpers below
func addressPbToModel(a *pb.Address) models.Address {
	if a == nil { return models.Address{} }
	return models.Address{ FullText: a.FullText, Lat: a.Lat, Lng: a.Lng }
}

func customerPbToModel(c *pb.CustomerSnapshot) models.CustomerSnapshot {
	if c == nil { return models.CustomerSnapshot{} }
	return models.CustomerSnapshot{ DisplayName: c.DisplayName, Phone: c.Phone }
}

func itemsPbToModel(items []*pb.WasteItem) []models.WasteItem {
	res := make([]models.WasteItem, 0, len(items))
	for _, it := range items {
		if it == nil { continue }
		res = append(res, models.WasteItem{ Type: it.Type, Weight: it.Weight })
	}
	return res
}

func statusFromString(s string) models.OrderStatus {
	switch s {
	case string(models.StatusCreated):
		return models.StatusCreated
	case string(models.StatusAccepted):
		return models.StatusAccepted
	case string(models.StatusOnWay):
		return models.StatusOnWay
	case string(models.StatusComplete):
		return models.StatusComplete
	case string(models.StatusCancelled):
		return models.StatusCancelled
	default:
		return models.StatusCreated
	}
}

func valueOrEmpty(p *string) string { if p == nil { return "" }; return *p }


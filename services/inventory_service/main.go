package main

import (
	"context"
	"log"
	"net"

	pb "carthage/protos/inventory_service"
	handler "carthage/services/inventory_service/biz"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterInventoryServiceServer(s, NewInventoryService())

	reflection.Register(s)

	l, tcpErr := net.Listen("tcp", ":50053")
	if tcpErr != nil {
		log.Fatalf("cannot listen on tcp port :50053 :%v", tcpErr)
	}

	grpcErr := s.Serve(l)
	if grpcErr != nil {
		log.Fatalf("cannot serve grpc server: %v", grpcErr)
	}

}

type InventoryService struct {
	pb.UnimplementedInventoryServiceServer
	handler handler.InventoryOpsInterface
}

type InventoryServiceInterface interface {
	GetInventoryListing(req *pb.InventoryListingRequest) pb.InventoryListingResponse
	CreateInventory(req *pb.CreateInventoryRequest) pb.CreateInventoryResponse
	UpdateInventory(req *pb.UpdateInventoryRequest) pb.UpdateInventoryResponse
	BulkCreateInventory(req *pb.BulkCreateInventoryRequest) pb.BulkCreateInventoryResponse
}

func NewInventoryService() *InventoryService {
	return &InventoryService{
		UnimplementedInventoryServiceServer: pb.UnimplementedInventoryServiceServer{},
		handler:                             handler.NewInventoryOpsHandler(),
	}
}

func (s *InventoryService) GetInventoryListing(ctx context.Context, req *pb.InventoryListingRequest) (*pb.InventoryListingResponse, error) {
	res, err := s.handler.ListInventoryWithFilters(req.TenantId, handler.InventoryFilters{
		Status:     req.Status.String(),
		EntityType: req.EntityType.String(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.InventoryListingResponse{
		Data:       res,
		Pagination: &pb.Pagination{},
	}, nil
}

func (s *InventoryService) CreateInventory(ctx context.Context, req *pb.CreateInventoryRequest) (*pb.CreateInventoryResponse, error) {
	return &pb.CreateInventoryResponse{}, nil
}

func (s *InventoryService) UpdateInventory(ctx context.Context, req *pb.UpdateInventoryRequest) (*pb.UpdateInventoryResponse, error) {
	return &pb.UpdateInventoryResponse{}, nil
}

func (s *InventoryService) BulkCreateInventory(ctx context.Context, req *pb.BulkCreateInventoryRequest) (*pb.BulkCreateInventoryResponse, error) {
	return &pb.BulkCreateInventoryResponse{}, nil
}

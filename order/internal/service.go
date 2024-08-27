package internal

import (
	"log"
	"net"

	"go-microservices/order/proto"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type OrderServiceServer struct {
	proto.UnimplementedOrderServiceServer
	repo OrderRepository
}

func NewOrderServiceServer(db *gorm.DB) *grpc.Server {
	repo := NewOrderRepository(db)

	grpcServer := grpc.NewServer()

	orderServiceServer := &OrderServiceServer{
		repo: repo,
	}
	proto.RegisterOrderServiceServer(grpcServer, orderServiceServer)

	return grpcServer
}

func StartOrderServer(port string, db *gorm.DB) {
	grpcServer := NewOrderServiceServer(db)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	// Start serving requests
	log.Printf("Order Service is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

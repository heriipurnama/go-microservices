package internal

import (
	"log"
	"net"

	"go-microservices/product/proto"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ProductServiceServer struct {
	proto.UnimplementedProductServiceServer
	repo ProductRepository
}

func NewProductServiceServer(db *gorm.DB) *grpc.Server {
	repo := NewProductRepository(db)

	grpcServer := grpc.NewServer()

	productServiceServer := &ProductServiceServer{
		repo: repo,
	}
	proto.RegisterProductServiceServer(grpcServer, productServiceServer)

	return grpcServer
}

func StartProductServer(port string, db *gorm.DB) {
	grpcServer := NewProductServiceServer(db)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	log.Printf("Product Service is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

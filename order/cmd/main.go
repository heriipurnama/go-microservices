package main

import (
	"log"
	"net"

	"go-microservices/order/internal"
	"go-microservices/order/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=user password=password dbname=order port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	orderRepo := internal.NewOrderRepository(db)
	orderService := &internal.OrderService{Repo: orderRepo}

	grpcServer := grpc.NewServer()
	proto.RegisterOrderServiceServer(grpcServer, orderService)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen on port 50052: %v", err)
	}

	log.Println("Order Service is running on port 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}

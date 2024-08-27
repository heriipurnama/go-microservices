package main

import (
	"log"
	"net"

	"go-microservices/product/internal"
	"go-microservices/product/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=user password=password dbname=product port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	jwtService := internal.NewJWTService("your-secret-key")
	productRepo := internal.NewProductRepository(db)
	productService := internal.NewProductService(productRepo, jwtService)

	grpcServer := grpc.NewServer()
	proto.RegisterProductServiceServer(grpcServer, productService)

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen on port 50053: %v", err)
	}

	log.Println("Product Service is running on port 50053")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}

package main

import (
	"log"
	"net"

	"go-microservices/auth/internal"
	"go-microservices/auth/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// note: masih harcoded - next: bikin file config dan menggunakan .env
	dsn := "host=localhost user=user password=password dbname=auth port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	jwtService := internal.NewJWTService("your-secret-key")
	userRepo := internal.NewUserRepository(db)
	authService := &internal.AuthService{
		Repo: userRepo,
		JWT:  jwtService,
	}

	grpcServer := grpc.NewServer()
	proto.RegisterAuthServiceServer(grpcServer, authService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	log.Println("Auth Service is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}

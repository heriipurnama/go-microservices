package internal

import (
	"log"
	"net"

	"go-microservices/auth/proto"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type AuthServiceServer struct {
	proto.UnimplementedAuthServiceServer
	repo UserRepository
	jwt  JWTService
}

func NewAuthServiceServer(db *gorm.DB, secretKey string) *grpc.Server {
	repo := NewUserRepository(db)
	jwtService := NewJWTService(secretKey)

	grpcServer := grpc.NewServer()

	authServiceServer := &AuthServiceServer{
		repo: repo,
		jwt:  jwtService,
	}
	proto.RegisterAuthServiceServer(grpcServer, authServiceServer)

	return grpcServer
}

func StartAuthServer(port string, db *gorm.DB, secretKey string) {
	grpcServer := NewAuthServiceServer(db, secretKey)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	log.Printf("Auth Service is running on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

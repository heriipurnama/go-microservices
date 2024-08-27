package internal

import (
	"context"
	"errors"

	"go-microservices/auth/proto"
)

type AuthService struct {
	proto.UnimplementedAuthServiceServer
	Repo UserRepository
	JWT  JWTService
}

func (s *AuthService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.AuthResponse, error) {
	user := User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	if err := s.Repo.CreateUser(user); err != nil {
		return nil, err
	}

	token, err := s.JWT.GenerateToken(user.Username)
	if err != nil {
		return nil, err
	}

	return &proto.AuthResponse{Token: token}, nil
}

func (s *AuthService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.AuthResponse, error) {
	user, err := s.Repo.GetUserByUsername(req.Username)
	if err != nil || !user.CheckPassword(req.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.JWT.GenerateToken(user.Username)
	if err != nil {
		return nil, err
	}

	return &proto.AuthResponse{Token: token}, nil
}

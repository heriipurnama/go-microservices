package internal

import (
	"context"

	"go-microservices/product/proto"
)

type ProductService struct {
	proto.UnimplementedProductServiceServer
	repo ProductRepository
	jwt  JWTService
}

func NewProductService(repo ProductRepository, jwt JWTService) *ProductService {
	return &ProductService{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.ProductResponse, error) {
	product := Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	if err := s.repo.CreateProduct(product); err != nil {
		return nil, err
	}

	return &proto.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *proto.GetProductRequest) (*proto.ProductResponse, error) {
	product, err := s.repo.GetProductByID(req.ProductId)
	if err != nil {
		return nil, err
	}

	return &proto.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

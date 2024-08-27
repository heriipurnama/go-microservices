package internal

import (
	"context"
	"errors"

	"go-microservices/order/proto"

	"gorm.io/gorm"
)

type OrderService struct {
	proto.UnimplementedOrderServiceServer
	Repo OrderRepository
}

func (s *OrderService) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.OrderResponse, error) {
	order := Order{
		ProductID: req.ProductId,
		Quantity:  req.Quantity,
		UserID:    req.UserId,
	}

	if err := s.Repo.CreateOrder(order); err != nil {
		return nil, err
	}

	return &proto.OrderResponse{
		Id:        order.ID,
		ProductId: order.ProductID,
		Quantity:  order.Quantity,
		UserId:    order.UserID,
	}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *proto.GetOrderRequest) (*proto.OrderResponse, error) {
	order, err := s.Repo.GetOrderByID(req.OrderId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	return &proto.OrderResponse{
		Id:        order.ID,
		ProductId: order.ProductID,
		Quantity:  order.Quantity,
		UserId:    order.UserID,
	}, nil
}

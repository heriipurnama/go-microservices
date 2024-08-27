package internal

import (
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order Order) error
	GetOrderByID(orderID string) (Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order Order) error {
	return r.db.Create(&order).Error
}

func (r *orderRepository) GetOrderByID(orderID string) (Order, error) {
	var order Order
	if err := r.db.First(&order, "id = ?", orderID).Error; err != nil {
		return Order{}, err
	}
	return order, nil
}

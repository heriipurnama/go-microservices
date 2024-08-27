package internal

import (
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product Product) error
	GetProductByID(productID string) (Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product Product) error {
	return r.db.Create(&product).Error
}

func (r *productRepository) GetProductByID(productID string) (Product, error) {
	var product Product
	if err := r.db.First(&product, "id = ?", productID).Error; err != nil {
		return Product{}, err
	}
	return product, nil
}

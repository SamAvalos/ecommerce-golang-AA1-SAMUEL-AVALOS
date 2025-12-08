package repository

import "ecommerce/models"

type ProductRepository interface {
	List() []models.Product
	FindByID(id string) (*models.Product, error)
	UpdateStock(id string, qty int) error
	Seed(items []models.Product)
}

type CartRepository interface {
	Get(userID string) (models.Cart, error)
	AddItem(userID string, item models.CartItem) error
	Clear(userID string) error
}

type OrderRepository interface {
	Create(order models.Order) error
	List() ([]models.Order, error)
}

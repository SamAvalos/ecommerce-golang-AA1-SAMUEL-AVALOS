package service

import (
	"ecommerce/internal/repository"
	"ecommerce/models"
	"errors"
)

type CartService struct {
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewCartService(cartRepo repository.CartRepository, productRepo repository.ProductRepository) *CartService {
	return &CartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

// GET CART
func (s *CartService) GetCart(userID string) (models.Cart, error) {
	return s.cartRepo.Get(userID)
}

// ADD ITEM
func (s *CartService) AddItem(userID string, item models.CartItem) error {
	// Validar que el producto exista
	product, err := s.productRepo.FindByID(item.ProductID)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	// Validar stock
	if item.Qty > product.Stock {
		return errors.New("insufficient stock")
	}

	// Rellenar precio real del producto
	item.PriceCents = product.PriceCents

	// Guardar ítem en el carrito
	return s.cartRepo.AddItem(userID, item)
}

// CLEAR CART (para después del checkout)
func (s *CartService) Clear(userID string) error {
	return s.cartRepo.Clear(userID)
}

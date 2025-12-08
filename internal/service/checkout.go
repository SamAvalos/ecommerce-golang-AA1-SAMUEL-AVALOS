package service

import (
	"ecommerce/internal/repository"
	"ecommerce/models"
	"errors"
	"fmt"
	"time"
)

type CheckoutService struct {
	products repository.ProductRepository
	carts    repository.CartRepository
	orders   repository.OrderRepository
}

func NewCheckoutService(p repository.ProductRepository, c repository.CartRepository, o repository.OrderRepository) *CheckoutService {
	return &CheckoutService{products: p, carts: c, orders: o}
}

// Checkout
func (s *CheckoutService) Checkout(userID string) (*models.Order, error) {
	// 1. Obtener carrito
	cart, err := s.carts.Get(userID)
	if err != nil {
		return nil, err
	}
	if len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	// 2. Reservar stock
	type stockOp struct {
		productID string
		qty       int
	}
	var reservations []stockOp

	for _, item := range cart.Items {
		if err := s.products.UpdateStock(item.ProductID, item.Qty); err != nil {
			// Rollback de lo ya reservado
			for _, op := range reservations {

				s.products.UpdateStock(op.productID, -op.qty)
			}
			return nil, err
		}
		reservations = append(reservations, stockOp{item.ProductID, item.Qty})
	}

	// 3. Calcular total y crear orden
	var total int64
	orderItems := make([]models.OrderItem, len(cart.Items))

	for i, item := range cart.Items {
		total += int64(item.Qty) * item.PriceCents
		orderItems[i] = models.OrderItem(item)
	}

	order := models.Order{
		ID:         fmt.Sprintf("ord_%d", time.Now().UnixNano()),
		UserID:     userID,
		Items:      orderItems,
		TotalCents: total,
		CreatedAt:  time.Now(),
	}

	if err = s.orders.Create(order); err != nil {
		for _, item := range cart.Items {
			s.products.UpdateStock(item.ProductID, +item.Qty)
		}
		return nil, err
	}

	s.carts.Clear(userID)
	return &order, nil
}

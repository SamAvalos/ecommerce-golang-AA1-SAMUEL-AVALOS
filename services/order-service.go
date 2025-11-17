package services

import (
	"ecommerce/models"
	"fmt"
	"sync"
	"time"
)

type OrdersService struct {
	orders []models.Order
	ps     *ProductsService
	mu     sync.RWMutex
}

func NewOrdersService(ps *ProductsService) *OrdersService {
	return &OrdersService{
		ps: ps,
	}
}

func (os *OrdersService) List() []models.Order {
	os.mu.RLock()
	defer os.mu.RUnlock()
	return os.orders
}

func (os *OrdersService) Create(userID string, items []models.CartItem) models.Order {
	os.mu.Lock()
	defer os.mu.Unlock()

	var total int64
	orderItems := []models.OrderItem{}
	for _, it := range items {
		total += int64(it.Qty) * it.PriceCents
		orderItems = append(orderItems, models.OrderItem{
			ProductID:  it.ProductID,
			Qty:        it.Qty,
			PriceCents: it.PriceCents,
		})
	}

	order := models.Order{
		ID:         fmt.Sprintf("order_%d", time.Now().UnixNano()),
		UserID:     userID,
		Items:      orderItems,
		TotalCents: total,
		CreatedAt:  time.Now(),
	}

	os.orders = append(os.orders, order)
	return order
}

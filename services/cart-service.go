package services

import (
	"ecommerce/models"
	"sync"
)

type CartService struct {
	carts map[string]models.Cart
	mu    sync.RWMutex
}

func NewCartService() *CartService {
	return &CartService{
		carts: make(map[string]models.Cart),
	}
}

func (cs *CartService) GetCart(userID string) models.Cart {
	cs.mu.RLock()
	cart, exists := cs.carts[userID]
	cs.mu.RUnlock()
	if !exists {
		cart = models.Cart{UserID: userID, Items: []models.CartItem{}}
		cs.mu.Lock()
		cs.carts[userID] = cart
		cs.mu.Unlock()
	}
	return cart
}

func (cs *CartService) AddItem(userID string, item models.CartItem) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cart := cs.carts[userID]
	cart.Items = append(cart.Items, item)
	cs.carts[userID] = cart
}

func (cs *CartService) ClearCart(userID string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.carts[userID] = models.Cart{UserID: userID, Items: []models.CartItem{}}
}

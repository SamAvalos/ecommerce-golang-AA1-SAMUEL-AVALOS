package services

import (
	"ecommerce/models"
	"sync"
)

type ProductsService struct {
	products []models.Product
	mu       sync.RWMutex
}

func NewProductsService() *ProductsService {
	return &ProductsService{}
}

func (ps *ProductsService) Seed(items []models.Product) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.products = items
}

func (ps *ProductsService) List() []models.Product {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	return ps.products
}

func (ps *ProductsService) GetByID(id string) *models.Product {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	for i := range ps.products {
		if ps.products[i].ID == id {
			return &ps.products[i]
		}
	}
	return nil
}

func (ps *ProductsService) UpdateStock(id string, qty int) bool {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for i := range ps.products {
		if ps.products[i].ID == id {
			if ps.products[i].Stock < qty {
				return false
			}
			ps.products[i].Stock -= qty
			return true
		}
	}
	return false
}

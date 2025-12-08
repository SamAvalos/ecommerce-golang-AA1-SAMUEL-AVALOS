package repository

import (
	"ecommerce/models"
	"errors"
	"sync"
)

// === Product In-Memory Repository ===
// Demuestra uso de: slice como backing store
type productRepo struct {
	products []models.Product // slice privado
	mu       sync.RWMutex
}

func NewProductRepository() ProductRepository {
	return &productRepo{
		products: make([]models.Product, 0, 20), // capacidad inicial
	}
}

func (r *productRepo) List() []models.Product {
	r.mu.RLock()
	defer r.mu.RUnlock()
	// Devuelve copia para evitar modificaciones externas (encapsulación)
	cpy := make([]models.Product, len(r.products))
	copy(cpy, r.products)
	return cpy
}

func (r *productRepo) FindByID(id string) (*models.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for i := range r.products {
		if r.products[i].ID == id {
			return &r.products[i], nil
		}
	}
	return nil, errors.New("product not found")
}

func (r *productRepo) UpdateStock(id string, qty int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.products {
		if r.products[i].ID == id {
			if r.products[i].Stock < qty {
				return errors.New("insufficient stock")
			}
			r.products[i].Stock -= qty
			return nil
		}
	}
	return errors.New("product not found")
}

func (r *productRepo) Seed(items []models.Product) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products = append(r.products[:0], items...) // reutiliza backing array
}

// === Cart In-Memory Repository ===
// Demuestra uso de: map[string]struct + slice interno
type cartRepo struct {
	carts map[string]models.Cart // key: userID
	mu    sync.RWMutex
}

func NewCartRepository() CartRepository {
	return &cartRepo{
		carts: make(map[string]models.Cart),
	}
}

func (r *cartRepo) Get(userID string) (models.Cart, error) {
	r.mu.RLock()
	cart, exists := r.carts[userID]
	r.mu.RUnlock()

	if !exists {
		// Crear carrito vacío si no existe (lazy initialization)
		cart = models.Cart{UserID: userID, Items: []models.CartItem{}}
		r.mu.Lock()
		r.carts[userID] = cart
		r.mu.Unlock()
	}
	return cart, nil
}

func (r *cartRepo) AddItem(userID string, item models.CartItem) error {
	if item.Qty <= 0 {
		return errors.New("quantity must be positive")
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	cart := r.carts[userID]
	cart.Items = append(cart.Items, item)
	r.carts[userID] = cart
	return nil
}

func (r *cartRepo) Clear(userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.carts[userID] = models.Cart{UserID: userID, Items: []models.CartItem{}}
	return nil
}

// === Order In-Memory Repository ===
// Demuestra uso de slice + generación de ID único
type orderRepo struct {
	orders []models.Order
	mu     sync.RWMutex
}

func NewOrderRepository() OrderRepository {
	return &orderRepo{}
}

func (r *orderRepo) Create(order models.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders = append(r.orders, order)
	return nil
}

func (r *orderRepo) List() ([]models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	cpy := make([]models.Order, len(r.orders))
	copy(cpy, r.orders)
	return cpy, nil
}

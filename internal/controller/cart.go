package controller

import (
	"encoding/json"
	"net/http"

	"ecommerce/internal/service"
	"ecommerce/models"
)

type CartController struct {
	cartService    *service.CartService
	productService *service.ProductService
}

func NewCartController(cs *service.CartService, ps *service.ProductService) *CartController {
	return &CartController{
		cartService:    cs,
		productService: ps,
	}
}

// GET /api/cart → protegido con JWT
func (c *CartController) GetCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(string)
	if !ok || userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	cart, err := c.cartService.GetCart(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

// POST /api/cart → protegido con JWT
func (c *CartController) AddToCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(string)
	if !ok || userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Quantity <= 0 {
		req.Quantity = 1
	}

	// Validamos producto y stock
	product, err := c.productService.GetByID(req.ProductID)
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	if product.Stock < req.Quantity {
		http.Error(w, "insufficient stock", http.StatusConflict)
		return
	}

	err = c.cartService.AddItem(userID, models.CartItem{
		ProductID:  req.ProductID,
		Qty:        req.Quantity,
		PriceCents: product.PriceCents,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "added to cart"})
}

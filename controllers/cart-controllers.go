package controllers

import (
	"ecommerce/models"
	"ecommerce/services"
	"encoding/json"
	"net/http"
)

type CartController struct {
	cs *services.CartService
	ps *services.ProductsService
	os *services.OrdersService
}

func NewCartController(cs *services.CartService, ps *services.ProductsService, os *services.OrdersService) *CartController {
	return &CartController{cs: cs, ps: ps, os: os}
}

func (cc *CartController) GetCart(w http.ResponseWriter, r *http.Request, userID string) {
	w.Header().Set("Content-Type", "application/json")
	cart := cc.cs.GetCart(userID)
	json.NewEncoder(w).Encode(cart)
}

func (cc *CartController) AddToCart(w http.ResponseWriter, r *http.Request, userID string, productID string, qty int) {
	product := cc.ps.GetByID(productID)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if qty <= 0 || qty > product.Stock {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}
	cc.cs.AddItem(userID, models.CartItem{ProductID: productID, Qty: qty, PriceCents: product.PriceCents})
	w.WriteHeader(http.StatusOK)
}

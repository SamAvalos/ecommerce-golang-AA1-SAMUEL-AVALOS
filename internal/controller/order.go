package controller

import (
	"encoding/json"
	"net/http"

	"ecommerce/internal/repository"
)

type OrderController struct {
	repo repository.OrderRepository // ← interface, no puntero a interface
}

func NewOrderController(repo repository.OrderRepository) *OrderController {
	return &OrderController{repo: repo}
}

func (c *OrderController) List(w http.ResponseWriter, r *http.Request) {
	orders, err := c.repo.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

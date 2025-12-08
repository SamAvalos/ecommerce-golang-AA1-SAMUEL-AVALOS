package controller

import (
	"encoding/json"
	"net/http"

	"ecommerce/internal/service"
)

type CheckoutController struct {
	service *service.CheckoutService
}

func NewCheckoutController(s *service.CheckoutService) *CheckoutController {
	return &CheckoutController{service: s}
}

// POST /api/checkout → protegido con JWT
func (c *CheckoutController) Checkout(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("userID").(string)
	if !ok || userID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	order, err := c.service.Checkout(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

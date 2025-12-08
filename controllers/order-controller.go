package controllers

import (
	"ecommerce/services"
	"encoding/json"
	"net/http"
)

type OrdersController struct {
	os *services.OrdersService
}

func NewOrdersController(os *services.OrdersService) *OrdersController {
	return &OrdersController{os: os}
}

func (oc *OrdersController) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(oc.os.List())
}

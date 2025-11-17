package controllers

import (
	"ecommerce/services"
	"encoding/json"
	"net/http"
)

type ProductsController struct {
	ps *services.ProductsService
}

func NewProductsController(ps *services.ProductsService) *ProductsController {
	return &ProductsController{ps: ps}
}

func (pc *ProductsController) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pc.ps.List())
}

package controller

import (
	"encoding/json"
	"net/http"

	"ecommerce/internal/service"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(s *service.ProductService) *ProductController {
	return &ProductController{service: s}
}

func (pc *ProductController) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products := pc.service.List()
	json.NewEncoder(w).Encode(products)
}

package service

import (
	"ecommerce/internal/repository"
	"ecommerce/models"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) List() []models.Product {
	return s.repo.List()
}

func (s *ProductService) GetByID(id string) (*models.Product, error) {
	return s.repo.FindByID(id)
}

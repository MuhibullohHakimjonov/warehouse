package services

import (
	"warehouse/api/new/internal/models"
	"warehouse/api/new/internal/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *ProductService) GetProductByName(name string) (*models.Product, error) {
	return s.repo.GetProductByName(name)
}

func (s *ProductService) CreateShelf(shelf *models.Shelf) error {
	return s.repo.CreateShelf(shelf)
}

package service

import (
	"hexa/internal/core/domain"
	"hexa/internal/core/ports"
)

type productService struct {
	repo ports.ProductRepository // Dependency Injection via Port
}

// Constructor untuk membuat service baru
func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAllProducts() ([]domain.Product, error) {
	// Di sini Anda bisa menambahkan business logic (misal: validasi, hitung diskon, dll)
	return s.repo.FindAll()
}
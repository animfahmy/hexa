package repository

import (
	"hexa/internal/core/domain"
	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *gormRepository {
	return &gormRepository{db: db}
}

// Implementasi fungsi FindAll dari interface ProductRepository
func (r *gormRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	// GORM melakukan query ke table 'products'
	result := r.db.Find(&products)
	return products, result.Error
}
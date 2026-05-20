package repository

import (
	"hexa/internal/core/domain"
	"gorm.io/gorm"
)

type GormRepositoryStruct struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepositoryStruct {
	return &GormRepositoryStruct{db: db}
}

// Implementasi fungsi FindAll dari interface ProductRepository
func (r *GormRepositoryStruct) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	// GORM melakukan query ke table 'products'
	result := r.db.Find(&products)
	return products, result.Error
}
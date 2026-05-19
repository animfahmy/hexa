package service

import (
	"errors"
	"hexa/internal/core/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1. Membuat Mock Struct yang meniru Repository asli
type mockProductRepository struct {
	mockData []domain.Product
	mockErr  error
}

// 2. Implementasikan method FindAll agar struct ini sah menjadi ProductRepository
func (m *mockProductRepository) FindAll() ([]domain.Product, error) {
	return m.mockData, m.mockErr
}

// 3. Fungsi Uji Pertama: Skenario Sukses mengambil data
func TestGetAllProducts_Success(t *testing.T) {
	// Siapkan data tiruan yang seolah-olah berasal dari MySQL
	dummyProducts := []domain.Product{
		{ID: 1, Name: "Laptop ASUS", Price: 15000000},
		{ID: 2, Name: "Mouse Logitech", Price: 300000},
	}

	// Buat instance repo palsu dengan data sukses
	repoPalsu := &mockProductRepository{
		mockData: dummyProducts,
		mockErr:  nil,
	}

	// Suntikkan repo palsu ke Service asli (Dependency Injection)
	service := NewProductService(repoPalsu)

	// Jalankan fungsi yang ingin kita uji
	hasil, err := service.GetAllProducts()

	// ASSERTION: Cek apakah hasilnya sesuai ekspektasi kita
	assert.NoError(t, err)                 // Pastikan tidak ada error
	assert.Len(t, hasil, 2)                // Pastikan jumlah data ada 2
	assert.Equal(t, "Laptop ASUS", hasil[0].Name) // Pastikan data pertama benar
}

// 4. Fungsi Uji Kedua: Skenario Gagal (Database Down / Error)
func TestGetAllProducts_Failed(t *testing.T) {
	// Buat instance repo palsu yang disetting melempar error
	repoPalsu := &mockProductRepository{
		mockData: nil,
		mockErr:  errors.New("database connection lost"),
	}

	service := NewProductService(repoPalsu)

	hasil, err := service.GetAllProducts()

	// ASSERTION
	assert.Error(t, err)                           // Pastikan error muncul
	assert.Nil(t, hasil)                           // Pastikan datanya kosong (nil)
	assert.Equal(t, "database connection lost", err.Error()) // Pastikan pesan errornya pas
}
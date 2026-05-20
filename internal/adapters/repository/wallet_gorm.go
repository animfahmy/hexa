package repository

import (
	"hexa/internal/core/domain"
	"gorm.io/gorm"
)

// PERBAIKAN: Ubah nama struct agar tidak bentrok dengan product_gorm.go
type walletRepository struct {
	db *gorm.DB
}

// PERBAIKAN: Ubah nama constructor-nya
func NewWalletRepository(db *gorm.DB) *walletRepository {
	return &walletRepository{db: db}
}

// PERBAIKAN: Sesuaikan dengan interface WalletRepository di ports.go
func (r *walletRepository) GetBalance(userID int) (int, error) {
	var wallet domain.Wallet
	
	// Query GORM untuk mengambil 1 data wallet berdasarkan UserID
	// SELECT * FROM wallets WHERE user_id = userID LIMIT 1;
	result := r.db.Where("user_id = ?", userID).First(&wallet)
	if result.Error != nil {
		return 0, result.Error
	}
	
	return wallet.Balance, nil
}
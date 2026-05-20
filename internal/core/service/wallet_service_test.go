package service

import (
	// "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1. Buat Mock Repository untuk Wallet
type mockWalletRepository struct {
	mockBalance int
	mockErr     error
}

func (m *mockWalletRepository) GetBalance(userID int) (int, error) {
	return m.mockBalance, m.mockErr
}

// 2. Test Skenario: Saldo Cukup
func TestCanWithdraw_Success(t *testing.T) {
	repoPalsu := &mockWalletRepository{
		mockBalance: 500000, // Set saldo palsu 500rb
		mockErr:     nil,
	}

	service := NewWalletService(repoPalsu)

	// Coba withdraw 200rb (harus lolos karena saldo 500rb)
	bisaWithdraw, err := service.CanWithdraw(1, 200000)

	assert.NoError(t, err)
	assert.True(t, bisaWithdraw)
}

// 3. Test Skenario: Saldo Kurang
func TestCanWithdraw_InsufficientBalance(t *testing.T) {
	repoPalsu := &mockWalletRepository{
		mockBalance: 500000, // Set saldo palsu 500rb
		mockErr:     nil,
	}

	service := NewWalletService(repoPalsu)

	// Coba withdraw 600rb (harus gagal karena saldo hanya 500rb)
	bisaWithdraw, err := service.CanWithdraw(1, 600000)

	assert.Error(t, err)
	assert.False(t, bisaWithdraw)
	assert.Equal(t, "saldo tidak mencukupi", err.Error())
}
package service

import (
	"errors"
	"hexa/internal/core/ports"
)

type walletService struct {
	repo ports.WalletRepository
}

func NewWalletService(repo ports.WalletRepository) ports.WalletService {
	return &walletService{repo: repo}
}

// PERBAIKAN: Parameter dan return value HARUS SAMA dengan yang ada di ports.go
func (s *walletService) CanWithdraw(userID int, amount int) (bool, error) {
	// 1. Ambil saldo dari repo (wajib masukkan userID sesuai kontrak)
	balance, err := s.repo.GetBalance(userID)
	if err != nil {
		return false, err
	}

	// 2. LOGIKA BISNIS: Cek apakah saldo cukup
	if balance < amount {
		return false, errors.New("saldo tidak mencukupi")
	}

	// 3. Jika cukup, kembalikan true
	return true, nil
}

func (s *walletService) GetBalance(userID int) (int, error) {
	return s.repo.GetBalance(userID)
}
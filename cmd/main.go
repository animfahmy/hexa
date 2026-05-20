package main

import (
	"hexa/internal/config" // Import package config yang baru dibuat
	"hexa/internal/adapters/handler"
	"hexa/internal/adapters/repository"
	"hexa/internal/core/service"

	"github.com/labstack/echo/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. Load semua konfigurasi dari .env
	cfg := config.LoadConfig()

	// 2. Gunakan DB DSN dari config struct
	db, err := gorm.Open(mysql.Open(cfg.DBString), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi ke database!")
	}

	e := echo.New()

	productRepo := repository.NewGormRepository(db)
	productSvc := service.NewProductService(productRepo)
	handler.NewHttpProductHandler(e, productSvc)

	walletRepo := repository.NewWalletRepository(db)
	walletSvc := service.NewWalletService(walletRepo)
	handler.NewHttpWalletHandler(e, walletSvc) // Daftarkan di sini

	// 3. Gunakan Port dari config struct
	e.Start(cfg.AppPort)
}
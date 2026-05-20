//go:build wireinject
// +build wireinject

package main

import (
	"hexa/internal/adapters/handler"
	"hexa/internal/adapters/repository"
	"hexa/internal/core/ports" // Tambahkan import ini untuk membaca interface
	"hexa/internal/core/service"

	"github.com/google/wire"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

// Penampung rakitan untuk Product
func InitializeProductHandler(e *echo.Echo, db *gorm.DB) *handler.HttpProductHandler {
	wire.Build(
		repository.NewGormRepository,
		// PERBAIKAN: Beri tahu Wire bahwa NewGormRepository (struct) ditujukan untuk mengisi ProductRepository (interface)
		wire.Bind(new(ports.ProductRepository), new(*repository.GormRepositoryStruct)), 
		service.NewProductService,
		handler.NewHttpProductHandlerReturn,
	)
	return nil
}

// Penampung rakitan untuk Wallet
func InitializeWalletHandler(e *echo.Echo, db *gorm.DB) *handler.HttpWalletHandler {
	wire.Build(
		repository.NewWalletRepository,
		// PERBAIKAN: Beri tahu Wire bahwa NewWalletRepository (struct) ditujukan untuk mengisi WalletRepository (interface)
		wire.Bind(new(ports.WalletRepository), new(*repository.WalletRepositoryStruct)), 
		service.NewWalletService,
		handler.NewHttpWalletHandlerReturn,
	)
	return nil
}
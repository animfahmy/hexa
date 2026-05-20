package handler

import (
	"net/http"
	"strconv" // Digunakan untuk konversi string ke int
	"hexa/internal/core/ports"
	"github.com/labstack/echo/v5"
)

type HttpWalletHandler struct {
	service ports.WalletService
}

func NewHttpWalletHandler(e *echo.Echo, svc ports.WalletService) {
	handler := &HttpWalletHandler{service: svc}
	
	// Route 1: Menampilkan saldo berdasarkan ID
	e.GET("/wallets/:id", handler.GetBalance)
	
	// Route 2: Mengecek kemampuan withdraw berdasarkan ID dan Nominal
	e.GET("/wallets/:id/withdraw/:nominal", handler.CheckWithdraw)
}

// Tambahkan fungsi baru ini di bawah NewHttpWalletHandler lama Anda:
func NewHttpWalletHandlerReturn(e *echo.Echo, svc ports.WalletService) *HttpWalletHandler {
	handler := &HttpWalletHandler{service: svc}
	e.GET("/wallets/:id", handler.GetBalance)
	e.GET("/wallets/:id/withdraw/:nominal", handler.CheckWithdraw)
	return handler
}

// 1. Handler untuk Cek Saldo
func (h *HttpWalletHandler) GetBalance(c *echo.Context) error {
	// Ambil id dari parameter URL lalu konversi string ke int
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID user harus berupa angka"})
	}

	balance, err := h.service.GetBalance(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Kembalikan response JSON yang rapi
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userID,
		"balance": balance,
	})
}

// 2. Handler untuk Cek Kemampuan Tarik Saldo
func (h *HttpWalletHandler) CheckWithdraw(c *echo.Context) error {
	// Ambil parameter id dan nominal dari URL
	userID, errId := strconv.Atoi(c.Param("id"))
	nominal, errNominal := strconv.Atoi(c.Param("nominal"))

	if errId != nil || errNominal != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID dan Nominal harus berupa angka"})
	}

	// Panggil logika bisnis di service
	bisaWithdraw, err := h.service.CanWithdraw(userID, nominal)
	if err != nil {
		// Jika saldo kurang, service akan melempar error "saldo tidak mencukupi"
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "insufficient balance",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "can withdraw",
		"user_id": userID,
		"amount":  nominal,
		"allowed": bisaWithdraw,
	})
}
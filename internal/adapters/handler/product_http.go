package handler

import (
	"net/http"
	"hexa/internal/core/ports"
	"github.com/labstack/echo/v5"
)

type HttpProductHandler struct {
	service ports.ProductService
}

func NewHttpProductHandler(e *echo.Echo, svc ports.ProductService) {
	handler := &HttpProductHandler{service: svc}
	
	// Daftarkan route Echo ke method handler
	e.GET("/products", handler.GetProducts)
}

// Tambahkan fungsi baru ini di bawah NewHttpProductHandler lama Anda:
func NewHttpProductHandlerReturn(e *echo.Echo, svc ports.ProductService) *HttpProductHandler {
	handler := &HttpProductHandler{service: svc}
	e.GET("/products", handler.GetProducts)
	return handler
}

func (h *HttpProductHandler) GetProducts(c *echo.Context) error {
	products, err := h.service.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}
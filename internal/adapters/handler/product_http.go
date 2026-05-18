package handler

import (
	"net/http"
	"hexa/internal/core/ports"
	"github.com/labstack/echo/v4"
)

type HttpProductHandler struct {
	service ports.ProductService
}

func NewHttpProductHandler(e *echo.Echo, svc ports.ProductService) {
	handler := &HttpProductHandler{service: svc}
	
	// Daftarkan route Echo ke method handler
	e.GET("/products", handler.GetProducts)
}

func (h *HttpProductHandler) GetProducts(c echo.Context) error {
	products, err := h.service.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}
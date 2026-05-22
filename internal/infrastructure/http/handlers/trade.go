package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"trade-organization/internal/application/trading"
	"trade-organization/internal/infrastructure/http/dto"
)

type TradeHandler struct {
	service *trading.TradeService
}

func NewTradeHandler(service *trading.TradeService) *TradeHandler {
	return &TradeHandler{service: service}
}

// CreateSale обрабатывает POST /api/v1/trade/sales
func (h *TradeHandler) CreateSale(c *gin.Context) {
	sellerID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user context not found"})
		return
	}

	var req dto.CreateSaleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sale, details := req.ToDomain(sellerID.(int))

	saleID, err := h.service.CreateSale(c.Request.Context(), sale, details)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"sale_id": saleID, "status": "success"})
}

// TransferProduct обрабатывает POST /api/v1/trade/transfer
func (h *TradeHandler) TransferProduct(c *gin.Context) {
	var req dto.TransferProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.TransferProduct(c.Request.Context(), req.FromStoreID, req.ToStoreID, req.ProductID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "transfer completed"})
}

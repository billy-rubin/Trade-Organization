package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"trade-organization/internal/application/supply"
	"trade-organization/internal/infrastructure/http/dto"
)

type SupplyHandler struct {
	service *supply.SupplyService
}

func NewSupplyHandler(service *supply.SupplyService) *SupplyHandler {
	return &SupplyHandler{service: service}
}

// CreateRequest обрабатывает POST /api/v1/supply/requests
func (h *SupplyHandler) CreateRequest(c *gin.Context) {
	var req dto.CreateSupplyRequestPayload

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq, details := req.ToDomain()

	requestID, err := h.service.CreateRequest(c.Request.Context(), domainReq, details)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"request_id": requestID,
		"status":     "request created successfully",
	})
}

// GenerateOrder обрабатывает POST /api/v1/supply/orders/generate
func (h *SupplyHandler) GenerateOrder(c *gin.Context) {
	var req dto.GenerateOrderPayload

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.ProcessRequestToOrder(c.Request.Context(), req.RequestID, req.SupplierID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "order generated successfully"})
}

// ReceiveOrder обрабатывает POST /api/v1/supply/orders/receive
func (h *SupplyHandler) ReceiveOrder(c *gin.Context) {
	var req dto.ReceiveOrderPayload

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.ReceiveOrder(c.Request.Context(), req.OrderID, req.StoreID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "order received and inventory updated"})
}

func (h *SupplyHandler) GetOrderDetails(c *gin.Context) {
	orderIDStr := c.Query("order_id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil || orderID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing order_id parameter"})
		return
	}

	details, err := h.service.GetOrderDetails(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch order details: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, details)
}

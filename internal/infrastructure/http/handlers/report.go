package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"trade-organization/internal/application/reports"
	"trade-organization/internal/infrastructure/http/dto"
)

type ReportHandler struct {
	service *reports.ReportService
}

func NewReportHandler(service *reports.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetInventory(c *gin.Context) {
	storeType := c.Query("store_type")
	if storeType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_type query parameter is required"})
		return
	}

	reportData, err := h.service.GetInventory(c.Request.Context(), storeType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToInventoryReportResponseList(reportData))
}

// GetProfitability обрабатывает запрос на рентабельность
func (h *ReportHandler) GetProfitability(c *gin.Context) {
	storeIDStr := c.Query("store_id")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid store_id"})
		return
	}

	startDate, err := time.Parse(time.DateOnly, c.Query("start_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format, use YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse(time.DateOnly, c.Query("end_date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format, use YYYY-MM-DD"})
		return
	}

	reportData, err := h.service.GetProfitability(c.Request.Context(), storeID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToProfitabilityResponse(reportData))
}

// GetTurnover обрабатывает запрос товарооборота. Пример: GET /api/v1/reports/turnover
func (h *ReportHandler) GetTurnover(c *gin.Context) {
	reportData, err := h.service.GetTurnover(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToTurnoverResponseList(reportData))
}

// GetStoreEfficiency — Обработчик для Запроса 7 (Эффективность точки)
func (h *ReportHandler) GetStoreEfficiency(c *gin.Context) {
	storeIDStr := c.Query("store_id")
	storeID, err := strconv.Atoi(storeIDStr)
	if err != nil || storeID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing store_id parameter"})
		return
	}

	report, err := h.service.GetStoreEfficiency(c.Request.Context(), storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch store efficiency: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

// GetSupplierDeliveries — Обработчик для Запроса 6 (Поставки поставщика)
func (h *ReportHandler) GetSupplierDeliveries(c *gin.Context) {
	supplierID, err1 := strconv.Atoi(c.Query("supplier_id"))
	productID, err2 := strconv.Atoi(c.Query("product_id"))
	if err1 != nil || err2 != nil || supplierID <= 0 || productID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid supplier_id or product_id"})
		return
	}

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, err3 := time.Parse("2006-01-02", startDateStr)
	endDate, err4 := time.Parse("2006-01-02", endDateStr)
	if err3 != nil || err4 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dates, use YYYY-MM-DD format"})
		return
	}

	deliveries, err := h.service.GetSupplierDeliveries(c.Request.Context(), supplierID, productID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch supplier deliveries: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, deliveries)
}

// GetProductCustomers — Обработчик для Запроса 10 (Покупатели товара в точке)
func (h *ReportHandler) GetProductCustomers(c *gin.Context) {
	productID, err1 := strconv.Atoi(c.Query("product_id"))
	storeID, err2 := strconv.Atoi(c.Query("store_id"))
	if err1 != nil || err2 != nil || productID <= 0 || storeID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product_id or store_id"})
		return
	}

	customers, err := h.service.GetProductCustomers(c.Request.Context(), productID, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch product customers: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, customers)
}

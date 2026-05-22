package handlers

import (
	"net/http"

	"trade-organization/internal/application/management" // Исправленный импорт
	"trade-organization/internal/infrastructure/http/dto"

	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	service *management.StoreService
}

func NewStoreHandler(service *management.StoreService) *StoreHandler {
	return &StoreHandler{service: service}
}

func (h *StoreHandler) CreateStore(c *gin.Context) {
	var req dto.CreateStoreRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainStore := req.ToDomain()

	storeID, err := h.service.CreateStore(c.Request.Context(), domainStore)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create store"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"store_id": storeID, "status": "created"})
}

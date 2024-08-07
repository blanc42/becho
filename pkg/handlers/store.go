package handlers

import (
	"net/http"
	"strconv"

	"github.com/blanc42/becho/pkg/handlers/request"
	"github.com/blanc42/becho/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type StoreHandler interface {
	CreateStore(c *gin.Context)
	GetStore(c *gin.Context)
	UpdateStore(c *gin.Context)
	DeleteStore(c *gin.Context)
	ListStores(c *gin.Context)
}

type storeHandler struct {
	storeUsecase usecase.StoreUsecase
}

func NewStoreHandler(storeUsecase usecase.StoreUsecase) StoreHandler {
	return &storeHandler{storeUsecase: storeUsecase}
}

func (h *storeHandler) CreateStore(c *gin.Context) {
	var req request.CreateStoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	store, err := h.storeUsecase.CreateStore(c, userID.(string), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    store,
		"message": nil,
		"error":   nil,
	})
}

func (h *storeHandler) GetStore(c *gin.Context) {
	store_id := c.Param("store_id")

	store, err := h.storeUsecase.GetStore(c, store_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": store})
}

func (h *storeHandler) UpdateStore(c *gin.Context) {
	id := c.Param("id")

	var req request.UpdateStoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedStore, err := h.storeUsecase.UpdateStore(c, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedStore)
}

func (h *storeHandler) DeleteStore(c *gin.Context) {
	id := c.Param("id")

	err := h.storeUsecase.DeleteStore(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store deleted successfully"})
}

func (h *storeHandler) ListStores(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.ParseInt(offsetStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	stores, err := h.storeUsecase.ListStores(c, int32(limit), int32(offset))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stores)
}

package handlers

import (
	"net/http"

	"github.com/blanc42/becho/apps/admin/pkg/handlers/request"
	"github.com/blanc42/becho/apps/admin/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type VariantHandler interface {
	CreateVariant(c *gin.Context)
	GetVariant(c *gin.Context)
	UpdateVariant(c *gin.Context)
	DeleteVariant(c *gin.Context)
	ListVariants(c *gin.Context)
}

type variantHandler struct {
	variantUsecase usecase.VariantUsecase
}

func NewVariantHandler(variantUsecase usecase.VariantUsecase) VariantHandler {
	return &variantHandler{variantUsecase: variantUsecase}
}

func (h *variantHandler) CreateVariant(c *gin.Context) {
	var req request.CreateVariantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storeID := c.Param("store_id")
	if storeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

	req.StoreID = storeID

	variantID, err := h.variantUsecase.CreateVariant(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"variant_id": variantID})
}

func (h *variantHandler) GetVariant(c *gin.Context) {
	id := c.Param("id")
	storeID := c.Value("store_id")
	if storeID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

	variant, err := h.variantUsecase.GetVariant(c, id, storeID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, variant)
}

func (h *variantHandler) UpdateVariant(c *gin.Context) {
	id := c.Param("id")
	var req request.UpdateVariantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storeID := c.GetString("store_id")
	if storeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

	req.StoreID = storeID

	updatedVariant, err := h.variantUsecase.UpdateVariant(c, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedVariant)
}

func (h *variantHandler) DeleteVariant(c *gin.Context) {
	id := c.Param("id")

	// TODO: get store_id from context
	storeID := "store_id"

	err := h.variantUsecase.DeleteVariant(c, id, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Variant deleted successfully"})
}

func (h *variantHandler) ListVariants(c *gin.Context) {
	storeID := c.Param("store_id")
	if storeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

	variants, err := h.variantUsecase.ListVariants(c, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, variants)
}

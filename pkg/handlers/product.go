package handlers

import (
	"net/http"
	"strconv"

	"github.com/blanc42/becho/pkg/handlers/request"
	"github.com/blanc42/becho/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	CreateProduct(c *gin.Context)
	GetProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	GetProducts(c *gin.Context)
}

type productHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) ProductHandler {
	return &productHandler{productUsecase: productUsecase}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var req request.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productID, err := h.productUsecase.CreateProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product_id": productID})
}

func (h *productHandler) GetProduct(c *gin.Context) {
	product_id := c.Param("product_id")
	store_id := c.Param("store_id")

	product, err := h.productUsecase.GetProduct(c, product_id, store_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var req request.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.productUsecase.UpdateProduct(c, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *productHandler) GetProducts(c *gin.Context) {
	storeID := c.Query("store_id")
	if storeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

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

	products, err := h.productUsecase.GetProducts(c, storeID, int32(limit), int32(offset))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

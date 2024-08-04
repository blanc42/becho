package handlers

import (
	"net/http"

	"github.com/blanc42/becho/pkg/handlers/request"
	"github.com/blanc42/becho/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	CreateCategory(c *gin.Context)
	GetCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	ListCategories(c *gin.Context)
}

type categoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) CategoryHandler {
	return &categoryHandler{categoryUsecase: categoryUsecase}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var req request.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryID, err := h.categoryUsecase.CreateCategory(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"category_id": categoryID})
}

func (h *categoryHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")

	category, err := h.categoryUsecase.GetCategory(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var req request.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := h.categoryUsecase.UpdateCategory(c, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	// TODO: get store_id from context
	storeID := "store_id"

	err := h.categoryUsecase.DeleteCategory(c, id, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (h *categoryHandler) ListCategories(c *gin.Context) {
	storeID := c.Query("store_id")
	if storeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

	categories, err := h.categoryUsecase.ListCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

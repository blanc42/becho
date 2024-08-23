package handlers

import (
	"net/http"

	"github.com/blanc42/becho/apps/admin/pkg/handlers/request"
	"github.com/blanc42/becho/apps/admin/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	CreateCategory(c *gin.Context)
	GetCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	ListCategories(c *gin.Context)
	GetAllCategories(c *gin.Context)
}

type categoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) CategoryHandler {
	return &categoryHandler{categoryUsecase: categoryUsecase}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	storeID := c.Param("store_id")
	if storeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

	var req request.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.StoreID = storeID

	categoryID, err := h.categoryUsecase.CreateCategory(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"category_id": categoryID})
}

func (h *categoryHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	storeID := c.Param("store_id")

	category, err := h.categoryUsecase.GetCategory(c, id, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	storeID := c.Param("store_id")
	var req request.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := h.categoryUsecase.UpdateCategory(c, id, storeID, req)
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
	storeID := c.Param("store_id")
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

func (h *categoryHandler) GetAllCategories(c *gin.Context) {
	storeID := c.Param("store_id")
	if storeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "store_id is required"})
		return
	}

	categories, err := h.categoryUsecase.GetAllCategoriesRecursive(c, storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

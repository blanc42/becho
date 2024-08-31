package handlers

import (
	"net/http"

	"github.com/blanc42/becho/apps/admin/pkg/handlers/request"
	"github.com/blanc42/becho/apps/admin/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type CartHandler interface {
	AddToCart(c *gin.Context)
	RemoveFromCart(c *gin.Context)
	UpdateCartItem(c *gin.Context)
	GetCart(c *gin.Context)
}

type cartHandler struct {
	cartUsecase usecase.CartUsecase
}

func NewCartHandler(cartUsecase usecase.CartUsecase) CartHandler {
	return &cartHandler{cartUsecase: cartUsecase}
}

func (h *cartHandler) AddToCart(c *gin.Context) {
	var req request.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cartResponse, err := h.cartUsecase.AddToCart(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartResponse)
}

func (h *cartHandler) RemoveFromCart(c *gin.Context) {
	cartItemID := c.Param("cart_item_id")

	err := h.cartUsecase.RemoveFromCart(c, cartItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}

func (h *cartHandler) UpdateCartItem(c *gin.Context) {
	var req request.UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cartResponse, err := h.cartUsecase.UpdateCartItem(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartResponse)
}

func (h *cartHandler) GetCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	cartResponse, err := h.cartUsecase.GetCart(c, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartResponse)
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/blanc42/becho/apps/admin/pkg/handlers/request"
	"github.com/blanc42/becho/apps/admin/pkg/usecase"
	"github.com/blanc42/becho/apps/admin/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateAdminUser(c *gin.Context)
	CreateCustomer(c *gin.Context)
	AdminLogin(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	Logout(c *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (h *userHandler) CreateAdminUser(c *gin.Context) {
	var req request.CreateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := h.userUsecase.CreateAdminUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_id": userID, "message": "Admin user created successfully"})
}

func (h *userHandler) CreateCustomer(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := h.userUsecase.CreateUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user_id": userID, "message": "Customer created successfully"})
}

func (h *userHandler) AdminLogin(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUsecase.AuthenticateUser(c, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateAccessAndRefreshTokens(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SetTokenCookie(c, accessToken)
	utils.SetRefreshTokenCookie(c, refreshToken)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "data": user})
}

func (h *userHandler) Logout(c *gin.Context) {
	utils.ClearTokenCookie(c)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (h *userHandler) GetUser(c *gin.Context) {
	userID := c.Value("user_id")
	role := c.Value("user_role")

	fmt.Println(userID, role)

	if role == "admin" {
		user, err := h.userUsecase.GetAdminUser(c, userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user, "message": "User retrieved successfully"})
	} else if role == "customer" {
		user, err := h.userUsecase.GetCustomer(c, userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": user, "message": "User retrieved successfully"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	}
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var req request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.userUsecase.UpdateUser(c, userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	err := h.userUsecase.DeleteUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

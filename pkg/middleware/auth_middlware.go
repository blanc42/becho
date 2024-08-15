package middleware

import (
	"net/http"

	db "github.com/blanc42/becho/pkg/db/sqlc"
	"github.com/blanc42/becho/pkg/domain"
	"github.com/blanc42/becho/pkg/initializers"
	"github.com/blanc42/becho/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("auth_token")
		refreshToken, _ := c.Cookie("refresh_token")

		if token == "" && refreshToken == "" {
			// Case 1: No auth token and no refresh token
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication token is required"})
			c.Abort()
			return
		}

		if token == "" && refreshToken != "" {
			// Case 2: No auth token but have refresh token
			newToken, err := utils.RefreshToken(refreshToken)
			if err != nil {
				utils.ClearTokenCookie(c)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
				c.Abort()
				return
			}

			newClaims, err := utils.ValidateToken(newToken)
			if err != nil {
				utils.ClearTokenCookie(c)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				c.Abort()
				return
			}

			userRepository := domain.NewUserRepository(db.NewDbStore(initializers.DbConnection))
			_, err = userRepository.GetUser(c, newClaims.UserID)
			if err != nil {
				utils.ClearTokenCookie(c)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
				c.Abort()
				return
			}

			utils.SetTokenCookie(c, newToken)
			c.Set("user_id", newClaims.UserID)
			c.Set("user_role", newClaims.Role)
		} else {
			// Case 3: Have auth token
			claims, err := utils.ValidateToken(token)
			if err == nil {
				// Token is valid
				c.Set("user_id", claims.UserID)
				c.Set("user_role", claims.Role)
				if claims.Role == "customer" {
					c.Set("store_id", claims.StoreID)
				}
			} else if refreshToken != "" {
				// Token is invalid but we have a refresh token
				newToken, err := utils.RefreshToken(refreshToken)
				if err != nil {
					utils.ClearTokenCookie(c)
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
					c.Abort()
					return
				}

				newClaims, err := utils.ValidateToken(newToken)
				if err != nil {
					utils.ClearTokenCookie(c)
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
					c.Abort()
					return
				}

				utils.SetTokenCookie(c, newToken)
				c.Set("user_id", newClaims.UserID)
				c.Set("user_role", newClaims.Role)
				if newClaims.Role == "customer" {
					c.Set("store_id", newClaims.StoreID)
				}
			} else {
				// Token is invalid and no refresh token
				utils.ClearTokenCookie(c)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found"})
			c.Abort()
			return
		}

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

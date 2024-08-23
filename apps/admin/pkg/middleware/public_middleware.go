package middleware

import (
	db "github.com/blanc42/becho/apps/admin/pkg/db/sqlc"
	"github.com/blanc42/becho/apps/admin/pkg/domain"
	"github.com/blanc42/becho/apps/admin/pkg/initializers"
	"github.com/blanc42/becho/apps/admin/pkg/utils"
	"github.com/gin-gonic/gin"
)

func PublicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := c.Cookie("auth_token")
		refreshToken, _ := c.Cookie("refresh_token")

		if token == "" && refreshToken == "" {
			c.Next()
			return
		}

		if token == "" && refreshToken != "" {
			// Case 2: No auth token but have refresh token
			newToken, err := utils.RefreshToken(refreshToken)
			if err != nil {
				utils.ClearTokenCookie(c)
				c.Next()
				return
			}

			newClaims, err := utils.ValidateToken(newToken)
			if err != nil {
				utils.ClearTokenCookie(c)
				c.Next()
				return
			}

			userRepository := domain.NewUserRepository(db.NewDbStore(initializers.DbConnection))
			_, err = userRepository.GetUser(c, newClaims.UserID)
			if err != nil {
				utils.ClearTokenCookie(c)
				c.Next()
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
					c.Next()
					return
				}

				newClaims, err := utils.ValidateToken(newToken)
				if err != nil {
					utils.ClearTokenCookie(c)
					c.Next()
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
				c.Next()
				return
			}
		}

		c.Next()
	}
}

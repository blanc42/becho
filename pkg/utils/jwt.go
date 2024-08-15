package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key") // Replace with a secure secret key

type Claims struct {
	UserID  string `json:"user_id"`
	Role    string `json:"role"`
	StoreID string `json:"store_id,omitempty"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token
func GenerateToken(userID, role, storeID string, expirationTime time.Time) (string, error) {
	claims := &Claims{
		UserID:  userID,
		Role:    role,
		StoreID: storeID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken checks if the token is valid and returns the claims
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshToken generates a new token with a new expiration time
func RefreshToken(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	// Set new expiration time
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims.ExpiresAt = jwt.NewNumericDate(expirationTime)

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return newToken.SignedString(jwtSecret)
}

// GetUserIDFromToken extracts the user ID from the token
func GetUserIDFromToken(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.UserID, nil
}

// GetRoleFromToken extracts the role from the token
func GetRoleFromToken(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.Role, nil
}

// GetStoreIDFromToken extracts the store ID from the token
func GetStoreIDFromToken(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.StoreID, nil
}

// IsTokenExpired checks if the token has expired
func IsTokenExpired(tokenString string) bool {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return true
	}
	return claims.ExpiresAt.Before(time.Now())
}

// GenerateAccessAndRefreshTokens creates both access and refresh tokens
func GenerateAccessAndRefreshTokens(userID, role string) (string, string, error) {
	accessTokenExpiration := time.Now().Add(5 * time.Minute)
	refreshTokenExpiration := time.Now().Add(7 * 24 * time.Hour)

	accessToken, err := GenerateToken(userID, role, "", accessTokenExpiration)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateToken(userID, role, "", refreshTokenExpiration)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// SetTokenCookie sets the JWT token in an HTTP-only cookie
func SetTokenCookie(c *gin.Context, token string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true, // Set to true if using HTTPS
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   int(time.Hour * 24 * 7 / time.Second), // 1 week
	})
}

func SetRefreshTokenCookie(c *gin.Context, token string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   int(time.Hour * 24 * 7 / time.Second),
	})
}

// ClearTokenCookie clears the auth_token cookie
func ClearTokenCookie(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   -1,
	})
}

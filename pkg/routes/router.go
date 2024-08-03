package routes

import (
	"net/http"

	"github.com/blanc42/becho/pkg/handlers"
	"github.com/gin-gonic/gin"
)

// gin router
func SetupRouter(e *gin.Engine, u handlers.UserHandler) {
	// publick routes
	r := e.Group("/api/v1")

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/signup", u.CreateAdminUser)

	// private routes

	// admin routes

}

package routes

import (
	"github.com/blanc42/becho/pkg/handlers"
	"github.com/blanc42/becho/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// gin router
func SetupRouter(e *gin.Engine, u handlers.UserHandler, p handlers.ProductHandler, s handlers.StoreHandler, c handlers.CategoryHandler, v handlers.VariantHandler) {
	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	e.Use(cors.New(config))

	// API v1 group
	api := e.Group("/api/v1")
	{
		// admin login router
		api.POST("/signup", u.CreateAdminUser)
		api.POST("/login", u.Login)

		api.GET("/stores/:store_id", s.GetStore)
		stores := api.Group("/stores/:store_id")
		{
			stores.GET("/login", u.Login)
			stores.POST("/signup", u.CreateCustomer)

			products := stores.Group("/products")
			{
				products.GET("/:product_id", p.GetProduct)
			}

			categories := stores.Group("/categories")
			{
				categories.GET("/:category_id", c.GetCategory)
				categories.GET("/", c.ListCategories)
			}

			variants := categories.Group("/:category_id/variants")
			{
				variants.GET("/:variant_id", v.GetVariant)
				variants.GET("/", v.ListVariants)
			}
		}

	}

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		stores := protected.Group("/stores/:store_id")
		stores.GET("/users/:user_id", u.GetUser)
		stores.PUT("/users/:user_id", u.UpdateUser)

		// TODO : cart routes
		// TODO : order routes

	}

	// Admin routes
	admin := protected.Group("/")
	admin.Use(middleware.AdminMiddleware())
	{
		stores := admin.Group("/stores")
		{
			stores.POST("/", s.CreateStore)
			stores.PUT("/:store_id", s.UpdateStore)
			stores.DELETE("/:store_id", s.DeleteStore)
			stores.GET("/", s.ListStores)

			categories := stores.Group("/:store_id/categories")
			{
				categories.POST("/", c.CreateCategory)
				categories.PUT("/:category_id", c.UpdateCategory)
				categories.DELETE("/:category_id", c.DeleteCategory)
			}

			variants := categories.Group("/:category_id/variants")
			{
				variants.POST("/", v.CreateVariant)
				variants.PUT("/:category_id", v.UpdateVariant)
				variants.DELETE("/:category_id", v.DeleteVariant)
			}

		}

	}

}

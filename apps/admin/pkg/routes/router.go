package routes

import (
	"github.com/blanc42/becho/apps/admin/pkg/handlers"
	"github.com/blanc42/becho/apps/admin/pkg/middleware"
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

	e.RedirectTrailingSlash = true

	e.Use(middleware.PublicMiddleware())
	api := e.Group("/api/v1")
	{
		// admin login router
		api.POST("/signup", u.CreateAdminUser)
		api.POST("/login", u.AdminLogin)

		api.GET("/stores/:store_id", s.GetStore)
		stores := api.Group("/stores/:store_id")
		{
			// stores.GET("/login", u.Login)
			stores.POST("/signup", u.CreateCustomer)

			stores.GET("/products", p.GetProducts)
			products := stores.Group("/products")
			{
				products.GET("/:product_id", p.GetProduct)
				products.POST("/", p.CreateProduct)
			}

			categories := stores.Group("/categories")
			{
				categories.GET("/", c.GetAllCategories)
				categories.GET("/:category_id", c.GetCategory)
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
	admin.GET("/user", u.GetUser)
	admin.POST("/logout", u.Logout)
	admin.POST("/stores", s.CreateStore)
	admin.PUT("/stores/:store_id", s.UpdateStore)
	admin.DELETE("/stores/:store_id", s.DeleteStore)
	admin.GET("/stores", s.ListStores)
	admin.POST("/stores/:store_id/categories", c.CreateCategory)
	admin.PUT("/stores/:store_id/categories/:category_id", c.UpdateCategory)
	admin.DELETE("/stores/:store_id/categories/:category_id", c.DeleteCategory)
	admin.GET("/stores/:store_id/variants", v.ListVariants)
	admin.POST("/stores/:store_id/variants", v.CreateVariant)
	admin.PUT("/stores/:store_id/variants/:variant_id", v.UpdateVariant)
	admin.DELETE("/stores/:store_id/variants/:variant_id", v.DeleteVariant)

}

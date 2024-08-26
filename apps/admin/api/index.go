package handler

import (
	"context"
	"net/http"

	db "github.com/blanc42/becho/apps/admin/pkg/db/sqlc"
	"github.com/blanc42/becho/apps/admin/pkg/domain"
	"github.com/blanc42/becho/apps/admin/pkg/handlers"
	"github.com/blanc42/becho/apps/admin/pkg/initializers"
	"github.com/blanc42/becho/apps/admin/pkg/routes"
	"github.com/blanc42/becho/apps/admin/pkg/services"
	"github.com/blanc42/becho/apps/admin/pkg/usecase"
	"github.com/gin-gonic/gin"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	initializers.InitDatabase(ctx)
	defer initializers.DbConnection.Close(ctx)

	router := gin.Default()

	if initializers.DbConnection == nil {
		panic("DB connection is not initialized")
	}
	userRepository := domain.NewUserRepository(db.NewDbStore(initializers.DbConnection))
	storeRepository := domain.NewStoreRepository(db.NewDbStore(initializers.DbConnection))
	categoryRepository := domain.NewCategoryRepository(db.NewDbStore(initializers.DbConnection))
	variantRepository := domain.NewVariantRepository(db.NewDbStore(initializers.DbConnection))
	productRepository := domain.NewProductRepository(db.NewDbStore(initializers.DbConnection))
	imageRepository := domain.NewImageRepository(db.NewDbStore(initializers.DbConnection))

	userHandler := handlers.NewUserHandler(usecase.NewUserUseCase(userRepository, storeRepository))
	storeHandler := handlers.NewStoreHandler(usecase.NewStoreUseCase(storeRepository, userRepository))
	categoryHandler := handlers.NewCategoryHandler(usecase.NewCategoryUseCase(categoryRepository, storeRepository, variantRepository))
	variantHandler := handlers.NewVariantHandler(usecase.NewVariantUseCase(variantRepository, storeRepository, imageRepository))
	productHandler := handlers.NewProductHandler(usecase.NewProductUseCase(productRepository, storeRepository, variantRepository, imageRepository))

	routes.SetupRouter(router, userHandler, productHandler, storeHandler, categoryHandler, variantHandler)

	router.GET("/api/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	router.GET("/api/v1/images/upload", func(c *gin.Context) {
		c.JSON(http.StatusOK, services.GetUploadcareSignedParams())
	})

	router.ServeHTTP(w, r)
}

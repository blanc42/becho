package handler

import (
	"context"
	"net/http"

	db "github.com/blanc42/becho/pkg/db/sqlc"
	"github.com/blanc42/becho/pkg/domain"
	"github.com/blanc42/becho/pkg/handlers"
	"github.com/blanc42/becho/pkg/initializers"
	"github.com/blanc42/becho/pkg/routes"
	"github.com/blanc42/becho/pkg/usecase"
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

	userHandler := handlers.NewUserHandler(usecase.NewUserUseCase(userRepository))
	storeHandler := handlers.NewStoreHandler(usecase.NewStoreUseCase(storeRepository, userRepository))
	categoryHandler := handlers.NewCategoryHandler(usecase.NewCategoryUseCase(categoryRepository, storeRepository))
	variantHandler := handlers.NewVariantHandler(usecase.NewVariantUseCase(variantRepository, storeRepository))
	productHandler := handlers.NewProductHandler(usecase.NewProductUseCase(productRepository, storeRepository))

	routes.SetupRouter(router, userHandler, productHandler, storeHandler, categoryHandler, variantHandler)

	router.ServeHTTP(w, r)
}

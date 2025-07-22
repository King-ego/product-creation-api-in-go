package routes

import (
	"api/controller"
	"api/repository"
	"api/usecase"
	"database/sql"

	"api/middlewares"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
	router *gin.Engine
	db     *sql.DB
}

func NewProductRouter(router *gin.Engine, db *sql.DB) ProductRouter {
	return ProductRouter{
		router: router,
		db:     db,
	}
}

func (ru *ProductRouter) Routers() {

	ProductRepository := repository.NewProductRepository(ru.db)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(ProductUseCase)

	ru.router.GET("/products", middlewares.MidJwt, ProductController.GetProducts)
	ru.router.POST("/products", ProductController.CreateProduct)
	ru.router.GET("/products/:id", ProductController.GetProduct)
}

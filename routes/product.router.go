package routes

import (
	"api/controller"
	"api/db"
	"api/repository"
	"api/usecase"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
	router *gin.Engine
}

func NewProductRouter(router *gin.Engine) ProductRouter {
	return ProductRouter{
		router: router,
	}
}

func (ru *ProductRouter) Routers() {
	dbConnect, error := db.ConnectDB()

	if error != nil {
		panic(error)
	}

	ProductRepository := repository.NewProductRepository(dbConnect)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	ProductController := controller.NewProductController(ProductUseCase)

	ru.router.GET("/products", ProductController.GetProducts)
	ru.router.POST("/product", ProductController.GetProducts)
}

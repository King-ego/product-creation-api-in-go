package main

import (
	"api/controller"
	"api/db"
	"api/repository"
	"api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnect, error := db.ConnectDB()

	if error != nil {
		panic(error)
	}

	ProductRepository := repository.NewProductRepository(dbConnect)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.POST("/product", ProductController.GetProducts)

	server.Run(":8000")
}

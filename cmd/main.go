package main

import (
	"api/controller"
	"api/db"
	"api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnect, error := db.ConnectDB()

	if error != nil {
		panic(error)
	}

	ProductUseCase := usecase.NewProductUseCase()
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}

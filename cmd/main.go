package main

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProductController := controller.NewProductController()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}

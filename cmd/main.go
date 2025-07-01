package main

import (
	"api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProductRoutes := routes.NewProductRouter(server)

	ProductRoutes.Routers()

	UserRoutes := routes.NewUserRouter(server)

	UserRoutes.Routers()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	server.Run(":8090")
}

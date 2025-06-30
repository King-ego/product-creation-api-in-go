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
	UserRepository := repository.NewUserRepository(dbConnect)

	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	UserUsecase := usecase.NewUserUseCase(UserRepository)

	ProductController := controller.NewProductController(ProductUseCase)
	UserController := controller.NewUserController(UserUsecase)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.POST("/product", ProductController.GetProducts)

	server.GET("/users", UserController.GetUsers)

	server.Run(":8090")
}

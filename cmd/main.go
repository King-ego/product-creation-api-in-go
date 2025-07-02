package main

import (
	"api/db"
	"api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnect, error := db.ConnectDB()

	if error != nil {
		panic(error)
	}

	Router := router.NewRouters(server, dbConnect)
	Router.Routers()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	server.Run(":8090")
}

package main

import (
	"api_go_rest/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// layer of controller
	productController := controller.NewProductController()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/product", productController.GetProducts)

	server.Run(":8000")
}

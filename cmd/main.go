package main

import (
	"api_go_rest/controller"
	"api_go_rest/db"
	"api_go_rest/repository"
	"api_go_rest/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	//layer repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//layer usecase
	ProductUsecase := usecase.NewProductUseCase(ProductRepository)

	// layer of controller
	productController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/product", productController.GetProducts)

	server.Run(":8000")
}

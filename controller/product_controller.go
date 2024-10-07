package controller

import (
	"api_go_rest/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//useCase
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	product := []model.Product{
		{
			ID:    1,
			Name:  "Batata frita",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, product)
}

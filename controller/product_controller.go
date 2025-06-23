package controller

import (
	"api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

func NewProductController() ProductController {
	return ProductController{}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    "aas",
			Name:  "Diego",
			Price: 12,
		},
	}

	ctx.JSON(http.StatusOK, products)
}

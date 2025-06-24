package controller

import (
	"api/model"
	"api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return ProductController{
		ProductUseCase: usecase,
	}
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

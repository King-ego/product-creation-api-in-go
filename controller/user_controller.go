package controller

import (
	"api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) UserController {
	return UserController{
		UserUsecase: usecase,
	}
}

func (p *UserController) GetUsers(ctx *gin.Context) {
	users, err := p.UserUsecase.GetUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, users)
}

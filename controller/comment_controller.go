package controller

import (
	"api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	CommentUseCase usecase.CommentUseCase
}

func NewCommentController(usecase usecase.CommentUseCase) CommentController {
	return CommentController{
		CommentUseCase: usecase,
	}
}

func (pu *CommentController) GetComments(ctx *gin.Context) {
	comments, err := pu.CommentUseCase.GetComments()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, comments)

}

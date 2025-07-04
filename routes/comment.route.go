package routes

import (
	"api/controller"
	"api/repository"
	"api/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type CommentRouters struct {
	router *gin.Engine
	db     *sql.DB
}

func NewCommentRouters(router *gin.Engine, db *sql.DB) CommentRouters {
	return CommentRouters{
		db:     db,
		router: router,
	}
}

func (pr *CommentRouters) Routers() {
	CommentRepository := repository.NewCommentRepository(pr.db)
	CommentUseCase := usecase.NewCommentUseCase(CommentRepository)
	CommentController := controller.NewCommentController(CommentUseCase)

	pr.router.GET("/comments", CommentController.GetComments)
}

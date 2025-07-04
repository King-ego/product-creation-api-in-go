package router

import (
	"api/routes"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	db     *sql.DB
}

func NewRouters(router *gin.Engine, db *sql.DB) Router {
	return Router{
		router: router,
		db:     db,
	}
}

func (ru *Router) Routers() {
	productRouter := routes.NewProductRouter(ru.router, ru.db)
	userRouter := routes.NewUserRouter(ru.router, ru.db)
	CommentRouter := routes.NewCommentRouters(ru.router, ru.db)

	productRouter.Routers()
	userRouter.Routers()
	CommentRouter.Routers()
}

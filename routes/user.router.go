package routes

import (
	"api/controller"
	"api/repository"
	"api/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	router *gin.Engine
	db     *sql.DB
}

func NewUserRouter(router *gin.Engine, db *sql.DB) UserRouter {
	return UserRouter{
		router: router,
	}
}

func (ru *UserRouter) Routers() {

	UserRepository := repository.NewUserRepository(ru.db)

	UserUseCase := usecase.NewUserUseCase(UserRepository)

	UserController := controller.NewUserController(UserUseCase)

	ru.router.GET("users", UserController.GetUsers)

}

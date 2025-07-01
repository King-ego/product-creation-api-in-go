package routes

import (
	"api/controller"
	"api/db"
	"api/repository"
	"api/usecase"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	router *gin.Engine
}

func NewUserRouter(router *gin.Engine) UserRouter {
	return UserRouter{
		router: router,
	}
}

func (ru *UserRouter) Routers() {
	dbConnect, error := db.ConnectDB()

	if error != nil {
		panic(error)
	}

	UserRepository := repository.NewUserRepository(dbConnect)

	UserUseCase := usecase.NewUserUseCase(UserRepository)

	UserController := controller.NewUserController(UserUseCase)

	ru.router.GET("users", UserController.GetUsers)

}

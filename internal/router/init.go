package router

import (
	"server/internal/api"
	"server/internal/repository"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

type routeStruct struct {
	app     *gin.Engine
	handler *api.ServiceServer
}

func NewRouter(app *gin.Engine, dao *repository.DAO) *routeStruct {
	exampleService := service.NewExampleService(dao)
	authService := service.NewAuthService(dao)
	objectService := service.NewObjectService(dao)
	handler := api.NewService(
		*dao,
		exampleService,
		authService,
		objectService,
	)
	return &routeStruct{
		app:     app,
		handler: handler,
	}
}

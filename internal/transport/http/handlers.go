package http2

import (
	"github.com/Ypxd/WebService/internal/service"
	v1 "github.com/Ypxd/WebService/internal/transport/http/v1"
	"github.com/Ypxd/WebService/utils"
	"github.com/gin-gonic/gin"
)

func NewHandlers(service *service.Service) *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
	)

	handlerV1 := v1.NewHandler(service)
	api := router.Group("api")
	{
		handlerV1.Init(api)
	}

	utils.NewRoutes(router)

	return router
}

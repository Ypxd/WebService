package v1

import (
	"github.com/Ypxd/WebService/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initItemRoutes(v1)
	}
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

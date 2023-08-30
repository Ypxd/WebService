package v1

import (
	"github.com/Ypxd/WebService/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initItemRoutes(api *gin.RouterGroup) {
	item := api.Group("/item")
	{
		item.GET("/get-items", h.getItemByID)
	}
}

func (h *Handler) getItemByID(c *gin.Context) {
	var (
		err error
		res []models.Items
	)

	ids := c.QueryArray("id")

	res, err = h.services.Item.GetItem(c.Request.Context(), ids)
	if err != nil {
		response(c, http.StatusInternalServerError, err, nil, nil)
		return
	}

	response(c, http.StatusOK, err, res, nil)
	return
}

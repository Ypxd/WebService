package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type route struct {
	Method  string `json:"method"`
	Path    string `json:"path"`
	Handler string `json:"handler"`
}

func getMetricsRoute(r *gin.Engine) {
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

func getHealthcheckRoute(r *gin.Engine) {
	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
}

func getRouteList(r *gin.Engine) {
	r.GET("/routes", func(c *gin.Context) {
		resp := make([]route, 0)
		for _, r := range r.Routes() {
			resp = append(resp, route{
				Method:  r.Method,
				Path:    r.Path,
				Handler: r.Handler,
			})
		}

		c.JSON(http.StatusOK, resp)
	})
}

func NewRoutes(r *gin.Engine) {
	getMetricsRoute(r)
	getHealthcheckRoute(r)
	getRouteList(r)
	return
}

package v1

import (
	"github.com/gin-gonic/gin"
	"log"
)

type httpResponse struct {
	ErrorText string      `json:"error_text"`
	HasError  bool        `json:"has_error"`
	Message   interface{} `json:"message"`
	Count     *int        `json:"count,omitempty"`
}

func response(c *gin.Context,
	statusCode int,
	err error,
	message interface{},
	count *int) {
	resp := httpResponse{
		Message: message,
		Count:   count,
	}
	if err != nil {
		log.Println(err.Error())
		resp.ErrorText = err.Error()
		resp.HasError = true
	}

	c.AbortWithStatusJSON(statusCode, resp)
}

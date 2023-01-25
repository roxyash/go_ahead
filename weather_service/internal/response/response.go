package response

import (
	"github.com/gin-gonic/gin"
	"github.com/roxyash/go_ahead/pkg/zaplogger"
)

type Response struct {
	logger zaplogger.Logger
}

func NewResponse(logger zaplogger.Logger) *Response {
	return &Response{
		logger: logger,
	}
}


// Generate error response method
func (r *Response) NewErrorResponse(c *gin.Context, statusCode int, message string) {
	r.logger.Errorf(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}

type ForecastResponse struct {
	Temp string `json:"temp"`
	City string `json:"city"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type AboutResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

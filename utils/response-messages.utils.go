package utils

import (
	"fmt"
	"net/http"

	"github.com/enockjeremi/queswer/formatter"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Success bool `json:"success"`
	Error   any  `json:"error"`
}

func NotFoundResponse(c *gin.Context, field string) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Success: false,
		Error: struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    404,
			Message: fmt.Sprintf("%v not found", field),
		},
	})
}

func BadRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		Success: false,
		Error: struct {
			Code    int               `json:"code"`
			Message map[string]string `json:"message"`
		}{
			Code:    400,
			Message: formatter.NewErrorFormatter().Formatter(err),
		},
	})
}

func ForbiddenResponse(c *gin.Context) {
	c.JSON(http.StatusForbidden, ErrorResponse{
		Success: false,
		Error: struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    403,
			Message: "operation not allowed",
		},
	})
}

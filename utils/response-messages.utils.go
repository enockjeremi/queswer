package utils

import (
	"net/http"

	"github.com/enockjeremi/queswer/formatter"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Success bool `json:"success"`
	Error   any  `json:"error"`
}

func NotFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, ErrorResponse{
		Success: false,
		Error: struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    404,
			Message: message,
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

func ForbiddenResponse(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, ErrorResponse{
		Success: false,
		Error: struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    403,
			Message: message,
		},
	})
}

func UnauthorizedResponse(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, ErrorResponse{
		Success: false,
		Error: struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    401,
			Message: message,
		},
	})
}

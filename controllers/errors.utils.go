package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundError(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"status":  "404",
		"message": msg,
	})
}

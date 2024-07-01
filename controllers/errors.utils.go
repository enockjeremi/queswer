package controllers

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandling(c *gin.Context, httpType int, msg string) {
	c.JSON(httpType, gin.H{
		"success": false,
		"message": msg,
	})
}

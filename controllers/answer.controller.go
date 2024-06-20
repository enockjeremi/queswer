package controllers

import (
	"net/http"

	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

func GetAllAnswer(c *gin.Context) {}
func PostAnswer(c *gin.Context) {
	var answer models.Answer
	if err := c.ShouldBindBodyWithJSON(&answer); err != nil {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
		return
	}
	err := services.CreateAnswer(&answer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, answer)
	}
}

func GetOneAnswer(c *gin.Context) {}
func PutAnswer(c *gin.Context)    {}
func DeleteAnswer(c *gin.Context) {}

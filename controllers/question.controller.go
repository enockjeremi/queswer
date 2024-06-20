package controllers

import (
	"net/http"

	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

func GetAllQuestion(c *gin.Context) {
	var question []models.Question
	if err := services.FindAllQuestion(&question); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, question)
}
func PostQuestion(c *gin.Context) {
	var question models.Question
	question.Answer = make([]models.Answer, 0)

	if err := c.ShouldBindBodyWithJSON(&question); err != nil {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
		return
	}
	err := services.CreateQuestion(&question)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, question)
	}
}

func GetOneQuestion(c *gin.Context) {}
func PutQuestion(c *gin.Context)    {}
func DeleteQuestion(c *gin.Context) {}

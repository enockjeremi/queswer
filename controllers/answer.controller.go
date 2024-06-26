package controllers

import (
	"net/http"
	"strconv"

	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

func toString(i uint) string {
	v := uint64(i)
	return strconv.FormatUint(uint64(v), 10)
}

func GetAllAnswer(c *gin.Context) {}

func PostAnswer(c *gin.Context) {
	var answer models.Answer
	var question models.Question

	if err := c.ShouldBindBodyWithJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	questionID := toString(answer.QuestionID)

	err := services.GetOneQuestion(&question, questionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "question not found",
		})
	} else {
		err := services.CreateAnswer(&answer)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusCreated, answer)
		}
	}

}

func GetOneAnswer(c *gin.Context) {}
func PutAnswer(c *gin.Context)    {}
func DeleteAnswer(c *gin.Context) {}

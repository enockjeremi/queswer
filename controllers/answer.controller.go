package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/enockjeremi/queswer/utils"
	"github.com/gin-gonic/gin"
)

func toString(i uint) string {
	v := uint64(i)
	return strconv.FormatUint(uint64(v), 10)
}

func GetAllAnswer(c *gin.Context) {
	answers := make([]models.Answer, 0)
	if err := services.FindAllAnswer(&answers); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, answers)
}

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
		utils.ErrorHandling(c, http.StatusNotFound, "question not found")
		return
	} else {
		err := services.CreateAnswer(&answer)
		if err != nil {
			utils.ErrorHandling(c, http.StatusNotFound, "Could not create answer")
			return
		} else {
			c.JSON(http.StatusCreated, answer)
		}
	}

}

func GetOneAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")
	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		utils.ErrorHandling(c, http.StatusNotFound, "answer not found")
		return
	}
	c.JSON(http.StatusOK, answer)
}
func PutAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")
	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		utils.ErrorHandling(c, http.StatusNotFound, "answer not found")
		return
	}
	c.BindJSON(&answer)

	err = services.UpdateAnswer(&answer, id)
	if err != nil {
		utils.ErrorHandling(c, http.StatusNotFound, fmt.Sprintf("Could not update question ID: %v", id))
		return
	} else {
		c.JSON(http.StatusOK, answer)
	}
}

func DeleteAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")
	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		utils.ErrorHandling(c, http.StatusNotFound, "answer not found")
		return
	}

	err = services.DeleteAnswer(&answer, id)
	if err != nil {
		utils.ErrorHandling(c, http.StatusNotFound, fmt.Sprintf("Could not delete question ID: %v", id))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "deleted successfully",
	})

}

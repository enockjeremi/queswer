package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enockjeremi/queswer/formatter"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
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
	c.JSON(http.StatusOK, &answers)
}

func PostAnswer(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	user := currentUser.(models.User)

	var answer models.Answer
	var question models.Question

	answer.User = user
	if err := c.ShouldBindBodyWithJSON(&answer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": formatter.NewErrorFormatter().Formatter(err)})
		return
	}
	questionID := toString(answer.QuestionID)

	err := services.GetOneQuestion(&question, questionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "question not found",
		})
		return
	} else {
		err := services.CreateAnswer(&answer)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "Could not create answer",
			})
			return
		} else {
			c.JSON(http.StatusCreated, &answer)
		}
	}

}

func GetOneAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")
	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "answer not found",
		})
		return
	}
	c.JSON(http.StatusOK, &answer)
}
func PutAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")
	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "answer not found",
		})
		return
	}
	c.BindJSON(&answer)

	err = services.UpdateAnswer(&answer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   fmt.Sprintf("Could not update answer ID: %v", id),
		})
		return
	} else {
		c.JSON(http.StatusOK, &answer)
	}
}

func DeleteAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")
	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "answer not found",
		})
		return
	}

	err = services.DeleteAnswer(&answer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   fmt.Sprintf("Could not delete answer ID: %v", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "deleted successfully",
	})

}

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

func GetOneAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")
	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		MessageNotFound(c, "answer not found")
	}
	c.JSON(http.StatusOK, answer)
}
func PutAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")
	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		MessageNotFound(c, "answer not found")
	}
	c.BindJSON(&answer)

	err = services.UpdateAnswer(&answer, id)
	if err != nil {
		MessageNotFound(c, "answer not found")
	} else {
		c.JSON(http.StatusOK, answer)
	}
}

func DeleteAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")
	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		MessageNotFound(c, "answer not found")
	}
	c.BindJSON(&answer)
	err = services.DeleteAnswer(&answer, id)
	if err != nil {
		MessageNotFound(c, "answer not found")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "deleted successfully",
		})
	}

}

func MessageNotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"status":  "404",
		"message": msg,
	})
}

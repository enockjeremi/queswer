package controllers

import (
	"fmt"
	"net/http"

	"github.com/enockjeremi/queswer/formatter"
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

	c.JSON(http.StatusOK, &question)
}

func PostQuestion(c *gin.Context) {
	var question models.Question
	currentUser, _ := c.Get("currentUser")
	user := currentUser.(models.User)

	question.Answer = make([]models.Answer, 0)
	question.User = user

	if err := c.ShouldBindBodyWithJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": formatter.NewErrorFormatter().Formatter(err)})
		return
	}
	err := services.CreateQuestion(&question)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "could not create question",
		})
		return
	} else {
		c.JSON(http.StatusCreated, question)
	}
}

func GetOneQuestion(c *gin.Context) {
	id := c.Params.ByName("id")
	var question models.Question
	err := services.GetOneQuestion(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "question not found",
		})
		return
	} else {
		c.JSON(http.StatusOK, question)
	}
}
func PutQuestion(c *gin.Context) {
	var question models.Question
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(models.User).ID

	err := services.GetOneQuestion(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "question not found",
		})
		return
	}

	if question.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "this question could not be updated",
		})
		return
	}
	c.BindJSON(&question)

	err = services.UpdateQuestion(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   fmt.Sprintf("Could not update question ID: %v", id),
		})
		return
	} else {
		c.JSON(http.StatusOK, question)
	}

}

func DeleteQuestion(c *gin.Context) {
	var question models.Question
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(models.User).ID

	err := services.GetOneQuestion(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "question not found",
		})
		return
	}

	if question.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "this question could not be deleted",
		})
		return
	}

	err = services.DeleteQuestion(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   fmt.Sprintf("could not delete question ID: %v", id),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messages": "deleted successfully",
		"success":  true,
	})

}

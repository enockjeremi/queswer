package controllers

import (
	"net/http"

	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/enockjeremi/queswer/utils"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func GetAllQuestion(c *gin.Context) {
	question := make([]models.Question, 0)
	if err := services.FindAllQuestion(&question); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &question,
	})
}

func PostQuestion(c *gin.Context) {
	var question models.Question
	currentUser, _ := c.Get("currentUser")
	user := currentUser.(models.User)

	question.Answer = make([]models.Answer, 0)
	question.User = user

	if err := c.ShouldBindBodyWithJSON(&question); err != nil {
		utils.BadRequestResponse(c, err)
		return
	}

	err := services.CreateQuestion(&question)
	if err != nil {
		utils.BadRequestResponse(c, err)
		return
	} else {
		c.JSON(http.StatusCreated, Response{
			Success: true,
			Data:    &question,
		})
	}
}

func GetOneQuestion(c *gin.Context) {
	id := c.Params.ByName("id")
	var question models.Question

	err := services.GetOneQuestion(&question, id)
	if err != nil {
		utils.NotFoundResponse(c, "question")
		return
	} else {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data:    &question,
		})
	}
}
func PutQuestion(c *gin.Context) {
	var question models.Question
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(models.User).ID

	err := services.GetOneQuestion(&question, id)
	if err != nil {
		utils.NotFoundResponse(c, "question")
		return
	}

	if question.UserID != userID {
		utils.ForbiddenResponse(c)
		return
	}
	c.BindJSON(&question)

	err = services.UpdateQuestion(&question)
	if err != nil {
		utils.NotFoundResponse(c, "update")
		return
	} else {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data:    &question,
		})
	}

}

func DeleteQuestion(c *gin.Context) {
	var question models.Question
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(models.User).ID

	err := services.GetOneQuestion(&question, id)
	if err != nil {
		utils.NotFoundResponse(c, "question")
		return
	}

	if question.UserID != userID {
		utils.ForbiddenResponse(c)
		return
	}

	err = services.DeleteQuestion(&question, id)
	if err != nil {
		utils.NotFoundResponse(c, "delete")
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &question,
	})

}

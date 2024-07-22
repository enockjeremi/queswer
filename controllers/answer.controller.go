package controllers

import (
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
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &answers,
	})
}

func PostAnswer(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	user := currentUser.(models.User)

	var answer models.Answer
	var question models.Question

	answer.User = user
	if err := c.ShouldBindBodyWithJSON(&answer); err != nil {
		utils.BadRequestResponse(c, err)
		return
	}
	questionID := toString(answer.QuestionID)

	err := services.GetOneQuestion(&question, questionID)
	if err != nil {
		utils.NotFoundResponse(c, "question")
		return
	} else {
		err := services.CreateAnswer(&answer)
		if err != nil {
			utils.BadRequestResponse(c, err)
			return
		} else {
			c.JSON(http.StatusCreated, Response{
				Success: true,
				Data:    &answer,
			})
		}
	}

}

func GetOneAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")
	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		utils.NotFoundResponse(c, "answer")
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &answer,
	})
}

func PutAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(models.User).ID

	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		utils.NotFoundResponse(c, "answer")
		return
	}
	c.BindJSON(&answer)

	if answer.UserID != userID {
		utils.ForbiddenResponse(c)
		return
	}

	err = services.UpdateAnswer(&answer)
	if err != nil {
		utils.NotFoundResponse(c, "update")
		return
	} else {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data:    &answer,
		})
	}
}

func DeleteAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(models.User).ID

	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		utils.NotFoundResponse(c, "answer")
		return
	}

	if answer.UserID != userID {
		utils.ForbiddenResponse(c)
		return
	}

	err = services.DeleteAnswer(&answer, id)
	if err != nil {
		utils.NotFoundResponse(c, "delete")
		return
	}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &answer,
	})

}

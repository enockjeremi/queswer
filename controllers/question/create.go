package question

import (
	"net/http"

	"github.com/enockjeremi/queswer/libs"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

func CreateQuestion(c *gin.Context) {
	var question models.Question
	currentUser, _ := c.Get("currentUser")
	user := currentUser.(models.User)

	question.Answer = make([]models.Answer, 0)
	question.User = user

	if err := c.ShouldBindBodyWithJSON(&question); err != nil {
		libs.BadRequestResponse(c, err)
		return
	}

	err := services.CreateQuestion(&question)
	if err != nil {
		libs.BadRequestResponse(c, err)
		return
	} else {
		c.JSON(http.StatusCreated, Response{
			Success: true,
			Data:    &question,
		})
	}
}

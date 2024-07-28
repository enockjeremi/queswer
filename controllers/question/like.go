package question

import (
	"net/http"

	"github.com/enockjeremi/queswer/libs"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

func LikeQuestion(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	questionId := c.Params.ByName("id")
	userId := currentUser.(models.User).ID
	var question models.Question

	err := services.GetOneQuestion(&question, questionId)
	if err != nil {
		libs.NotFoundResponse(c, "question not found")
		return
	}

	like := libs.ContainsInSlice(question.Likes, int64(userId))
	if !like {
		question.Likes = append(question.Likes, int64(userId))
	} else {
		question.Likes = libs.RemoveElement(question.Likes, int64(userId))
	}

	err = services.UpdateQuestion(&question)
	if err != nil {
		libs.NotFoundResponse(c, "could not update")
		return
	} else {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data:    &question,
		})
	}
}

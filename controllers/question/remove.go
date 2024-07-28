package question

import (
	"net/http"

	"github.com/enockjeremi/queswer/libs"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

func RemoveQuestion(c *gin.Context) {
	var question models.Question
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(models.User).ID

	err := services.GetOneQuestion(&question, id)
	if err != nil {
		libs.NotFoundResponse(c, "question not found")
		return
	}

	if question.UserID != userID {
		libs.ForbiddenResponse(c, "operation not allowed")
		return
	}

	err = services.DeleteQuestion(&question, id)
	if err != nil {
		libs.NotFoundResponse(c, "could not delete")
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &question,
	})

}

package answer

import (
	"net/http"

	"github.com/enockjeremi/queswer/libs"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

func UpdateAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")

	currentUser, _ := c.Get("currentUser")
	userID := currentUser.(models.User).ID

	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		libs.NotFoundResponse(c, "answer not found")
		return
	}
	c.BindJSON(&answer)

	if answer.UserID != userID {
		libs.ForbiddenResponse(c, "operation not allowed")
		return
	}

	err = services.UpdateAnswer(&answer)
	if err != nil {
		libs.NotFoundResponse(c, "could not update")
		return
	} else {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data:    &answer,
		})
	}
}

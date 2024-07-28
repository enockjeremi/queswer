package answer

import (
	"net/http"
	"strconv"

	"github.com/enockjeremi/queswer/libs"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

func toString(i uint) string {
	v := uint64(i)
	return strconv.FormatUint(uint64(v), 10)
}

func CreateAnswer(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	user := currentUser.(models.User)

	var answer models.Answer
	var question models.Question

	answer.User = user
	if err := c.ShouldBindBodyWithJSON(&answer); err != nil {
		libs.BadRequestResponse(c, err)
		return
	}
	questionID := toString(answer.QuestionID)

	err := services.GetOneQuestion(&question, questionID)
	if err != nil {
		libs.NotFoundResponse(c, "question not found")
		return
	} else {
		err := services.CreateAnswer(&answer)
		if err != nil {
			libs.BadRequestResponse(c, err)
			return
		} else {
			c.JSON(http.StatusCreated, Response{
				Success: true,
				Data:    &answer,
			})
		}
	}

}

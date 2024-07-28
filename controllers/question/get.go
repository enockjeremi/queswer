package question

import (
	"net/http"

	"github.com/enockjeremi/queswer/libs"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func GetOneQuestion(c *gin.Context) {
	id := c.Params.ByName("id")
	var question models.Question

	err := services.GetOneQuestion(&question, id)
	if err != nil {
		libs.NotFoundResponse(c, "question not found")
		return
	} else {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Data:    &question,
		})
	}
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

package answer

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

func GetOneAnswer(c *gin.Context) {
	var answer models.Answer
	id := c.Params.ByName("id")
	err := services.FindOneAnswer(&answer, id)
	if err != nil {
		libs.NotFoundResponse(c, "answer not found")
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    &answer,
	})
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiDescription struct {
	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

func SetupRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, ApiDescription{
			Title:       "Queswer",
			Version:     "1.0.0",
			Description: "Esta es una api para hacer preguntas y respuestas.",
		})
	})
	v1 := r.Group("/v1")

	QuestionRouter(v1)
	AnswerRouter(v1)
	AuthRouter(v1)

	return r
}

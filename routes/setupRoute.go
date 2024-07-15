package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":       "Queswer",
			"version":     "1.0.0",
			"description": "Esta es una api para hacer preguntas y respuestas.",
		})
	})
	v1 := r.Group("/v1")

	QuestionRouter(v1)
	AnswerRouter(v1)
	AuthRouter(v1)

	return r
}

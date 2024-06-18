package routes

import "github.com/gin-gonic/gin"

func SetupRoute() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")

	QuestionRouter(v1)
	AnswerRouter(v1)
	UserRoute(v1)

	return r
}

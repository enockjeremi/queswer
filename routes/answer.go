package routes

import (
	"github.com/enockjeremi/queswer/controllers/answer"
	"github.com/enockjeremi/queswer/middlewares"
	"github.com/gin-gonic/gin"
)

func AnswerRouter(g *gin.RouterGroup) {
	{
		g.GET("answer", answer.GetAllAnswer)
		g.GET("answer/:id", answer.GetOneAnswer)
		g.POST("answer", middlewares.CheckAuth, answer.CreateAnswer)
		g.PUT("answer/:id", middlewares.CheckAuth, answer.UpdateAnswer)
		g.DELETE("answer/:id", middlewares.CheckAuth, answer.RemoveAnswer)
	}
}

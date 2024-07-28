package routes

import (
	"github.com/enockjeremi/queswer/controllers/question"
	"github.com/enockjeremi/queswer/middlewares"
	"github.com/gin-gonic/gin"
)

func QuestionRouter(g *gin.RouterGroup) {
	{
		g.GET("question", question.GetAllQuestion)
		g.GET("question/:id", question.GetOneQuestion)
		g.POST("question", middlewares.CheckAuth, question.CreateQuestion)
		g.PUT("question/:id", middlewares.CheckAuth, question.UpdateQuestion)
		g.DELETE("question/:id", middlewares.CheckAuth, question.RemoveQuestion)
		g.PATCH("question/:id/like", middlewares.CheckAuth, question.LikeQuestion)
	}
}

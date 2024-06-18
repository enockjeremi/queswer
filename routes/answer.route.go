package routes

import (
	"github.com/enockjeremi/queswer/controllers"
	"github.com/gin-gonic/gin"
)

func AnswerRouter(g *gin.RouterGroup) {
	{
		g.GET("answer", controllers.GetAllAnswer)
		g.POST("answer", controllers.PostAnswer)
		g.GET("answer/:id", controllers.GetOneAnswer)
		g.PUT("answer/:id", controllers.PutAnswer)
		g.DELETE("answer/:id", controllers.DeleteAnswer)
	}
}

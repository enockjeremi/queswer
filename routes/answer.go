package routes

import (
	"github.com/enockjeremi/queswer/controllers"
	"github.com/enockjeremi/queswer/middlewares"
	"github.com/gin-gonic/gin"
)

func AnswerRouter(g *gin.RouterGroup) {
	{
		g.GET("answer", controllers.GetAllAnswer)
		g.GET("answer/:id", controllers.GetOneAnswer)
		g.POST("answer", middlewares.CheckAuth, controllers.PostAnswer)
		g.PUT("answer/:id", middlewares.CheckAuth, controllers.PutAnswer)
		g.DELETE("answer/:id", middlewares.CheckAuth, controllers.DeleteAnswer)
	}
}

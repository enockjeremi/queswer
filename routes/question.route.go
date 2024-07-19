package routes

import (
	"github.com/enockjeremi/queswer/controllers"
	"github.com/enockjeremi/queswer/middlewares"
	"github.com/gin-gonic/gin"
)

func QuestionRouter(g *gin.RouterGroup) {
	{
		g.GET("question", controllers.GetAllQuestion)
		g.GET("question/:id", controllers.GetOneQuestion)
		g.POST("question", middlewares.CheckAuth, controllers.PostQuestion)
		g.PUT("question/:id", middlewares.CheckAuth, controllers.PutQuestion)
		g.DELETE("question/:id", middlewares.CheckAuth, controllers.DeleteQuestion)
	}
}

package routes

import (
	"github.com/enockjeremi/queswer/controllers"
	"github.com/gin-gonic/gin"
)

func QuestionRouter(g *gin.RouterGroup) {
	{
		g.GET("question", controllers.GetAllQuestion)
		g.POST("question", controllers.PostQuestion)
		g.GET("question/:id", controllers.GetOneQuestion)
		g.PUT("question/:id", controllers.PutQuestion)
		g.DELETE("question/:id", controllers.DeleteQuestion)
	}
}

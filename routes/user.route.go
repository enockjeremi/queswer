package routes

import (
	"github.com/enockjeremi/queswer/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(g *gin.RouterGroup) {
	{
		g.GET("user", controllers.GetAllUser)
		g.POST("user", controllers.PostUser)
		g.GET("user/:id", controllers.GetOneUser)
		g.PUT("user/:id", controllers.PutUser)
		g.DELETE("user/:id", controllers.DeleteUser)
	}
}

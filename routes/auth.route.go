package routes

import (
	"github.com/enockjeremi/queswer/controllers"
	"github.com/enockjeremi/queswer/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoute(g *gin.RouterGroup) {
	{
		g.POST("auth/sign-in", controllers.SignIn)
		g.POST("auth/sign-up", controllers.SignUp)
		g.GET("auth/profile", middlewares.CheckAuth, controllers.GetProfile)
	}
}

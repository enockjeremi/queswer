package routes

import (
	"github.com/enockjeremi/queswer/controllers/auth"
	"github.com/enockjeremi/queswer/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRouter(g *gin.RouterGroup) {
	{
		g.POST("auth/sign-in", auth.SignIn)
		g.POST("auth/sign-up", auth.SignUp)
		g.GET("auth/profile", middlewares.CheckAuth, auth.GetProfile)
		g.PUT("auth/update-profile", middlewares.CheckAuth, auth.UpdateProfile)
		g.PATCH("auth/change-password", middlewares.CheckAuth, auth.ChangePassword)
	}
}

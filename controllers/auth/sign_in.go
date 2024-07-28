package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/enockjeremi/queswer/libs"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PasswordInput struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=7"`
}

func SignIn(c *gin.Context) {
	var signIn AuthInput
	var user models.User

	if err := c.ShouldBindBodyWithJSON(&signIn); err != nil {
		libs.BadRequestResponse(c, err)
		return
	}

	err := services.VerifyUsername(&user, signIn.Username)
	if err != nil {
		libs.ForbiddenResponse(c, "invalid username")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signIn.Password)); err != nil {
		libs.ForbiddenResponse(c, "invalid password")
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("jWT_SECRET")))
	if err != nil {
		libs.NotFoundResponse(c, "failed to generate token")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   token,
	})

}

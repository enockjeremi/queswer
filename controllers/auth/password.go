package auth

import (
	"net/http"

	"github.com/enockjeremi/queswer/libs"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(c *gin.Context) {
	var passwordInput PasswordInput
	var user models.User
	currentUser, _ := c.Get("currentUser")
	id := currentUser.(models.User).ID

	if err := c.ShouldBindBodyWithJSON(&passwordInput); err != nil {
		libs.BadRequestResponse(c, err)
		return
	}

	err := services.GetUser(&user, id)
	if err != nil {
		libs.NotFoundResponse(c, "user not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordInput.OldPassword)); err != nil {
		libs.NotFoundResponse(c, "old password not found")
		return
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(passwordInput.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		libs.NotFoundResponse(c, err.Error())
		return
	}

	err = services.ChangePassword(&user, string(newPasswordHash))
	if err != nil {
		libs.NotFoundResponse(c, "could not change password")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": "successful password change",
		"success":  true,
	})

}

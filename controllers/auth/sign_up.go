package auth

import (
	"net/http"

	"github.com/enockjeremi/queswer/libs"
	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var auth models.User

	if err := c.ShouldBindBodyWithJSON(&auth); err != nil {
		libs.BadRequestResponse(c, err)
		return
	}
	if err := services.VerifyCredentials(&auth); err == nil {
		libs.NotFoundResponse(c, "user or email already exists")
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		libs.NotFoundResponse(c, err.Error())
		return
	}

	profile := auth.Profile
	err = services.CreateProfile(&profile)
	if err != nil {
		libs.NotFoundResponse(c, "something wrong! could not register user")
		return
	}

	user := models.User{
		Username:  auth.Username,
		Email:     auth.Email,
		Password:  string(passwordHash),
		ProfileID: profile.ID,
		Profile:   profile,
	}

	err = services.CreateUser(&user)
	if err != nil {
		libs.NotFoundResponse(c, "something wrong! could not register user")
		return
	} else {
		c.JSON(http.StatusCreated, user)
	}

}

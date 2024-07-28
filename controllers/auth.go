package controllers

import (
	"fmt"
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

func UpdateProfile(c *gin.Context) {
	var profile models.Profile
	var user models.User
	currentUser, _ := c.Get("currentUser")
	profileId := currentUser.(models.User).ProfileID

	err := services.GetProfile(&profile, profileId)
	if err != nil {
		libs.NotFoundResponse(c, "profile not found")
		return
	}

	if err = c.ShouldBindBodyWithJSON(&profile); err != nil {
		libs.BadRequestResponse(c, err)
		return
	}

	err = services.UpdateProfile(&profile)
	if err != nil {
		libs.NotFoundResponse(c, fmt.Sprintf("Could not update user profile ID: %v", profileId))
		return
	}

	err = services.GetUser(&user, profile.UserID)
	if err != nil {
		libs.NotFoundResponse(c, "profile not found")
		return
	}
	c.JSON(http.StatusOK, user)

}

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

func GetProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})
}

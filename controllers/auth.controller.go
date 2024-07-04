package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/enockjeremi/queswer/models"
	"github.com/enockjeremi/queswer/services"
	"github.com/enockjeremi/queswer/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignIn(c *gin.Context) {
	var signIn AuthInput
	var user models.User

	if err := c.ShouldBindBodyWithJSON(&signIn); err != nil {
		utils.ErrorHandling(c, http.StatusBadRequest, err.Error())
		return
	}

	err := services.VerifyUsername(&user, signIn.Username)
	if err != nil {
		utils.ErrorHandling(c, http.StatusForbidden, "invalid username")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signIn.Password)); err != nil {
		utils.ErrorHandling(c, http.StatusForbidden, "invalid user or password")
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("jWT_SECRET")))
	if err != nil {
		utils.ErrorHandling(c, http.StatusBadRequest, "failed to generate token")
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
		utils.ErrorHandling(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := services.VerifyCredentials(&auth); err == nil {
		utils.ErrorHandling(c, http.StatusBadRequest, "user or email already exists")
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorHandling(c, http.StatusNotFound, err.Error())
		return
	}

	profile := auth.Profile
	err = services.CreateProfile(&profile)
	if err != nil {
		utils.ErrorHandling(c, 404, "Could not register user")
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
		utils.ErrorHandling(c, 404, "Could not register user")
		return
	} else {
		c.JSON(http.StatusCreated, user)
	}

}

func UpdateProfile(c *gin.Context) {
	var profile models.Profile
	var user models.User
	profileId, _ := c.Get("profileID")

	err := services.GetProfile(&profile, profileId)
	if err != nil {
		utils.ErrorHandling(c, http.StatusNotFound, "profile not found")
		return
	}

	if err = c.ShouldBindBodyWithJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.UpdateProfile(&profile)
	if err != nil {
		utils.ErrorHandling(c, http.StatusNotFound, fmt.Sprintf("Could not update user profile ID: %v", profileId))
		return
	}

	err = services.GetUser(&user, profile.UserID)
	if err != nil {
		utils.ErrorHandling(c, http.StatusNotFound, "profile not found")
		return
	}
	c.JSON(http.StatusOK, user)

}

func GetProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})
}

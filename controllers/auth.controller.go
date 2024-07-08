package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/enockjeremi/queswer/formatter"
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

func SignIn(c *gin.Context) {
	var signIn AuthInput
	var user models.User

	if err := c.ShouldBindBodyWithJSON(&signIn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": formatter.NewErrorFormatter().Formatter(err)})
		return
	}

	err := services.VerifyUsername(&user, signIn.Username)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "invalid user or password",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signIn.Password)); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "invalid user or password",
		})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("jWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "failed to generate token",
		})
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
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": formatter.NewErrorFormatter().Formatter(err)})
		return
	}
	if err := services.VerifyCredentials(&auth); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "user or email already exists",
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	profile := auth.Profile
	err = services.CreateProfile(&profile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Could not register user",
		})
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
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Could not register user",
		})
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
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "profile not found",
		})
		return
	}

	if err = c.ShouldBindBodyWithJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": formatter.NewErrorFormatter().Formatter(err)})
		return
	}

	err = services.UpdateProfile(&profile)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   fmt.Sprintf("Could not update user profile ID: %v", profileId),
		})
		return
	}

	err = services.GetUser(&user, profile.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "profile not found",
		})
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
